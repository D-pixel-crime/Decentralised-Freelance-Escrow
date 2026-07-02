"use client";

import React, { useCallback, useEffect, useState } from "react";
import { Loader2, Send, CheckCircle2, XCircle } from "lucide-react";
import { useEscrowAction } from "./useEscrowAction";
import type { Job, JobStatus } from "@/types/job";

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

  const requestPayment = useEscrowAction(
    job.contractAddress,
    "requestPayment",
    makeBusyCb("requestPayment")
  );

  const acceptJob = useEscrowAction(
    job.contractAddress,
    "acceptJobCompletion",
    makeBusyCb("acceptJob")
  );

  const rejectJob = useEscrowAction(
    job.contractAddress,
    "rejectJobCompletion",
    makeBusyCb("rejectJob")
  );

  const anyBusy = Object.values(busyMap).some(Boolean);

  // ── Freelancer: Request Payment ───────────────────────────────────
  if (role === "freelancer" && job.status === "ALL_STAKED_AND_PENDING") {
    return (
      <div className="mt-4 space-y-2">
        <ActionButton
          onClick={requestPayment.execute}
          disabled={anyBusy}
          isPending={requestPayment.isPending}
          isConfirming={requestPayment.isConfirming}
          isSuccess={requestPayment.isSuccess}
          icon={Send}
          label="Request Payment"
          successLabel="Payment requested!"
        />
        {requestPayment.error && (
          <p className="text-xs text-red-400/80 truncate" title={requestPayment.error.message}>
            {requestPayment.error.message.length > 80
              ? requestPayment.error.message.slice(0, 80) + "…"
              : requestPayment.error.message}
          </p>
        )}
      </div>
    );
  }

  // ── Client: Approve / Reject Work ─────────────────────────────────
  if (role === "client" && job.status === "PENDING_CLIENT_CONFIRMATION") {
    return (
      <div className="mt-4 space-y-2">
        <ActionButton
          onClick={acceptJob.execute}
          disabled={anyBusy}
          isPending={acceptJob.isPending}
          isConfirming={acceptJob.isConfirming}
          isSuccess={acceptJob.isSuccess}
          variant="success"
          icon={CheckCircle2}
          label="Approve Work"
          successLabel="Work approved — freelancer paid!"
        />
        {!acceptJob.isSuccess && (
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
        {(acceptJob.error || rejectJob.error) && (
          <p className="text-xs text-red-400/80 truncate">
            {(acceptJob.error || rejectJob.error)?.message?.slice(0, 80)}
          </p>
        )}
      </div>
    );
  }

  return null;
}
