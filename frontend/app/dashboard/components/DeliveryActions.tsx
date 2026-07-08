"use client";

import React, { useCallback, useEffect, useState } from "react";
import { Loader2, Send, CheckCircle2, XCircle, FileText, ExternalLink } from "lucide-react";
import axios from "axios";
import { useEscrowAction } from "./useEscrowAction";
// eslint-disable-next-line @typescript-eslint/no-unused-vars
import type { Job, JobStatus } from "@/types/job";
import { useWriteContract, useWaitForTransactionReceipt } from "wagmi";
import { FREELANCE_ESCROW_ABI } from "@/constants/contract";
import { parseEther } from "viem";
import { useQueryClient } from "@tanstack/react-query";
import { useToast } from "@/contexts/ToastContext";
import { extractErrorMsg } from "@/lib/utils";

// ── Button styling helpers ─────────────────────────────────────────────────

function ActionButton({
  onClick,
  disabled,
  isPending,
  isConfirming,
  isSuccess,
  variant = "primary",
  icon: Icon,
  label,
  pendingLabel = "Confirming in Wallet…",
  confirmingLabel = "Waiting for Chain…",
  successLabel = "Confirmed!",
}: {
  onClick: () => void;
  disabled: boolean;
  isPending: boolean;
  isConfirming: boolean;
  isSuccess: boolean;
  variant?: "primary" | "success" | "danger-ghost";
  icon: React.ComponentType<{ className?: string }>;
  label: string;
  pendingLabel?: string;
  confirmingLabel?: string;
  successLabel?: string;
}) {
  const isWaiting = isPending || isConfirming;

  const variantStyles: Record<string, { bg: string; glow: string; hoverBg: string; hoverGlow: string }> = {
    primary: {
      bg: `rgba(var(--vault-accent), 0.8)`,
      glow: `0 0 20px rgba(var(--vault-accent), 0.2)`,
      hoverBg: `rgba(var(--vault-accent), 0.9)`,
      hoverGlow: `0 0 28px rgba(var(--vault-accent), 0.3)`,
    },
    success: {
      bg: `rgba(16, 185, 129, 0.8)`,
      glow: `0 0 20px rgba(16, 185, 129, 0.2)`,
      hoverBg: `rgba(16, 185, 129, 0.9)`,
      hoverGlow: `0 0 28px rgba(16, 185, 129, 0.3)`,
    },
    "danger-ghost": {
      bg: `rgba(239, 68, 68, 0.12)`,
      glow: `none`,
      hoverBg: `rgba(239, 68, 68, 0.22)`,
      hoverGlow: `0 0 20px rgba(239, 68, 68, 0.15)`,
    },
  };

  const v = variantStyles[variant];

  if (isSuccess) {
    return (
      <div className="flex items-center gap-2 rounded-xl border border-emerald-500/20 bg-emerald-500/10 px-4 py-2.5">
        <span className="h-2 w-2 rounded-full bg-emerald-400" />
        <span className="text-xs font-medium text-emerald-400">
          {successLabel}
        </span>
      </div>
    );
  }

  return (
    <button
      onClick={onClick}
      disabled={disabled || isWaiting}
      className={`flex w-full items-center justify-center gap-2 rounded-xl px-4 py-2.5 text-sm font-semibold transition-all duration-200 disabled:cursor-not-allowed disabled:opacity-40 ${
        variant === "danger-ghost" ? "text-red-400 border border-red-500/20" : "text-white"
      }`}
      style={{
        background: v.bg,
        boxShadow: v.glow,
      }}
      onMouseEnter={(e) => {
        if (!isWaiting && !disabled) {
          e.currentTarget.style.background = v.hoverBg;
          e.currentTarget.style.boxShadow = v.hoverGlow;
        }
      }}
      onMouseLeave={(e) => {
        e.currentTarget.style.background = v.bg;
        e.currentTarget.style.boxShadow = v.glow;
      }}
    >
      {isWaiting ? (
        <>
          <Loader2 className="h-4 w-4 animate-spin" />
          {isPending ? pendingLabel : confirmingLabel}
        </>
      ) : (
        <>
          <Icon className="h-4 w-4" />
          {label}
        </>
      )}
    </button>
  );
}

// ── DeliveryActions ────────────────────────────────────────────────────────

export default function DeliveryActions({
  job,
  role,
  onBusyChange,
}: {
  job: Job;
  role: string;
  onBusyChange: (busy: boolean) => void;
}) {
  const toast = useToast();
  // Track aggregated busy state from multiple hooks
  const [busyMap, setBusyMap] = useState<Record<string, boolean>>({});

  const makeBusyCb = useCallback(
    (key: string) => (busy: boolean) => {
      setBusyMap((prev) => {
        if (prev[key] === busy) return prev;
        return { ...prev, [key]: busy };
      });
    },
    []
  );

  // Bubble up aggregated busy
  useEffect(() => {
    const anyBusy = Object.values(busyMap).some(Boolean);
    onBusyChange(anyBusy);
  }, [busyMap, onBusyChange]);

  const [showEvidenceUI, setShowEvidenceUI] = useState(false);
  const [isSubmittingEvidence, setIsSubmittingEvidence] = useState(false);
  const [evidenceFile, setEvidenceFile] = useState<File | null>(null);
  
  const [payAmount, setPayAmount] = useState<string>(job.payMin !== undefined && job.payMin > 0 ? job.payMin.toString() : "1");
  const queryClient = useQueryClient();

  const handleRequestPaymentFlow = async () => {
    if (!evidenceFile) return;
    setIsSubmittingEvidence(true);
    const setBusy = makeBusyCb("submitEvidence");
    setBusy(true);

    try {
      const uploadData = new FormData();
      uploadData.append("file", evidenceFile);
      const res = await axios.post("/api/ipfs/file", uploadData);
      
      if (!res.data?.IpfsHash) {
        throw new Error("Failed to get CID from IPFS proxy");
      }
      
      await axios.post(
        `${process.env.NEXT_PUBLIC_API_URL}/api/post/job/deliver`,
        {
          jobId: job.id,
          deliverableCid: res.data.IpfsHash
        },
        { withCredentials: true }
      );
      
      requestPayment.execute();
      setShowEvidenceUI(false);
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } catch (err: any) {
      toast.error("Failed to submit evidence: " + extractErrorMsg(err));
    } finally {
      setIsSubmittingEvidence(false);
      setBusy(false);
    }
  };

  const requestPayment = useEscrowAction(
    job.contractAddress,
    "requestPayment",
    makeBusyCb("requestPayment")
  );

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const { data: acceptHash, writeContract: acceptWriteContract, isPending: isAcceptPending, error: acceptError } = useWriteContract();
  const { isLoading: isAcceptConfirming, isSuccess: isAcceptSuccess } = useWaitForTransactionReceipt({ hash: acceptHash });

  // Bubble accept job busy state
  useEffect(() => {
    makeBusyCb("acceptJob")(isAcceptPending || isAcceptConfirming);
  }, [isAcceptPending, isAcceptConfirming, makeBusyCb]);

  useEffect(() => {
    if (isAcceptSuccess) {
      queryClient.invalidateQueries({ queryKey: ["myJobs"] });
    }
  }, [isAcceptSuccess, queryClient]);

  const handleAcceptWork = () => {
    acceptWriteContract({
      abi: FREELANCE_ESCROW_ABI,
      address: job.contractAddress as `0x${string}`,
      functionName: "acceptJobCompletion",
      value: parseEther(payAmount || "0"),
    }, {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      onError: (err: any) => {
        let msg = err.shortMessage || err.message || "Transaction failed";
        if (msg.toLowerCase().includes("user rejected")) {
          msg = "Transaction rejected by user.";
        }
        toast.error(msg);
      },
      onSuccess: () => {
        toast.success("Transaction submitted to network.");
      }
    });
  };

  const rejectJob = useEscrowAction(
    job.contractAddress,
    "rejectJobCompletion",
    makeBusyCb("rejectJob")
  );

  const anyBusy = Object.values(busyMap).some(Boolean);

  // ── Freelancer: Request Payment ───────────────────────────────────
  if (role === "freelancer" && job.status === "ALL_STAKED_AND_PENDING") {
    if (showEvidenceUI) {
      return (
        <div className="mt-4 space-y-4 rounded-xl border border-[rgba(var(--vault-accent),0.2)] bg-[rgba(var(--vault-accent),0.05)] p-4">
          <h4 className="text-sm font-semibold text-slate-200">Submit Work Evidence</h4>
          <p className="text-xs text-slate-400">Upload a PDF summarizing or containing your delivered work.</p>
          <input
            type="file"
            accept=".pdf"
            onChange={(e) => setEvidenceFile(e.target.files?.[0] || null)}
            disabled={anyBusy || isSubmittingEvidence}
            className="w-full text-sm text-slate-300 file:mr-4 file:rounded-lg file:border-0 file:bg-[rgba(var(--vault-accent),0.1)] file:px-4 file:py-2 file:text-sm file:font-semibold file:text-[rgba(var(--vault-accent),1)] hover:file:bg-[rgba(var(--vault-accent),0.2)]"
          />
          <div className="flex gap-2">
            <button
              onClick={() => setShowEvidenceUI(false)}
              disabled={anyBusy || isSubmittingEvidence}
              className="flex-1 rounded-xl border border-slate-700/60 bg-slate-800 px-4 py-2 text-sm font-medium text-slate-300 hover:bg-slate-700 disabled:opacity-50"
            >
              Cancel
            </button>
            <ActionButton
              onClick={handleRequestPaymentFlow}
              disabled={!evidenceFile || anyBusy}
              isPending={isSubmittingEvidence || requestPayment.isPending}
              isConfirming={requestPayment.isConfirming}
              isSuccess={requestPayment.isSuccess}
              icon={Send}
              label="Submit & Request"
              pendingLabel="Uploading Evidence..."
              successLabel="Payment requested!"
            />
          </div>
        </div>
      );
    }

    return (
      <div className="mt-4 space-y-2">
        <ActionButton
          onClick={() => setShowEvidenceUI(true)}
          disabled={anyBusy}
          isPending={false}
          isConfirming={false}
          isSuccess={requestPayment.isSuccess}
          icon={FileText}
          label="Submit Evidence & Request Payment"
          successLabel="Payment requested!"
        />
      </div>
    );
  }

  // ── Client: Approve / Reject Work ─────────────────────────────────
  if (role === "client" && job.status === "PENDING_CLIENT_CONFIRMATION") {
    return (
      <div className="mt-4 space-y-2">
        {job.deliverableCid && (
          <a
            href={`https://gateway.pinata.cloud/ipfs/${job.deliverableCid}`}
            target="_blank"
            rel="noopener noreferrer"
            className="flex w-full items-center justify-center gap-2 rounded-xl border border-[rgba(var(--vault-accent),0.4)] bg-[rgba(var(--vault-accent),0.1)] px-4 py-2.5 text-sm font-semibold text-[rgba(var(--vault-accent),1)] hover:bg-[rgba(var(--vault-accent),0.2)] transition-colors mb-4"
          >
            <ExternalLink className="h-4 w-4" />
            View Work Evidence
          </a>
        )}
        
        {!isAcceptSuccess && (
          <div className="flex flex-col gap-3 rounded-xl border border-slate-700/60 bg-slate-900/60 p-4">
            <p className="text-xs text-slate-400 italic">
              Note: The amount you have already staked will go to the freelancer. Enter the remaining/total amount you wish to pay accordingly.
            </p>
            <div className="flex items-center justify-between gap-3 bg-slate-800/40 rounded-lg px-3 py-2 border border-slate-700/40">
              <label className="text-xs font-semibold text-slate-300">Pay Amount (ETH):</label>
              <input 
                type="number" 
                step="0.01"
                min="0"
                value={payAmount}
                onChange={(e) => setPayAmount(e.target.value)}
                disabled={anyBusy}
                className="w-24 bg-transparent text-sm text-right text-slate-200 outline-none placeholder-slate-600 font-mono disabled:opacity-50"
              />
            </div>
          </div>
        )}

        <ActionButton
          onClick={handleAcceptWork}
          disabled={anyBusy}
          isPending={isAcceptPending}
          isConfirming={isAcceptConfirming}
          isSuccess={isAcceptSuccess}
          variant="success"
          icon={CheckCircle2}
          label="Approve Work & Pay"
          successLabel="Work approved — freelancer paid!"
        />
        {!isAcceptSuccess && (
          <ActionButton
            onClick={rejectJob.execute}
            disabled={anyBusy}
            isPending={rejectJob.isPending}
            isConfirming={rejectJob.isConfirming}
            isSuccess={rejectJob.isSuccess}
            variant="danger-ghost"
            icon={XCircle}
            label="Reject Work"
            successLabel="Work rejected — dispute raised."
          />
        )}
      </div>
    );
  }

  return null;
}
