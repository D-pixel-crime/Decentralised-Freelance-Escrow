"use client";

import React, { useCallback, useEffect, useRef, useState } from "react";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { useWriteContract, useWaitForTransactionReceipt, useAccount } from "wagmi";
import { parseEther } from "viem";
import type { Job, JobStatus } from "@/types/job";
import { FREELANCE_ESCROW_ABI } from "@/constants/contract";
import {
  Briefcase, ExternalLink, Loader2, Inbox, Wallet, Plus, X,
  Calendar, Mail, DollarSign, FileText, Type,
} from "lucide-react";
import AppNavbar from "@/components/app-navbar";
import { ObjectId } from "bson";
import JobCardActions from "./components/JobCardActions";
import FreelancerProfile from "./components/FreelancerProfile";
import ApplicantReviewModal from "./components/ApplicantReviewModal";
import { extractErrorMsg } from "@/lib/utils";

// ── Status badge styling ────────────────────────────────────────────────────

const STATUS_STYLES: Record<string, { bg: string; text: string; dot: string }> = {
  UNALLOCATED:                 { bg: "bg-slate-500/10",  text: "text-slate-400",  dot: "bg-slate-400"  },
  AGREED:                      { bg: "bg-blue-500/10",   text: "text-blue-400",   dot: "bg-blue-400"   },
  CLIENT_STAKED:               { bg: "bg-cyan-500/10",   text: "text-cyan-400",   dot: "bg-cyan-400"   },
  FREELANCER_STAKED:           { bg: "bg-teal-500/10",   text: "text-teal-400",   dot: "bg-teal-400"   },
  ALL_STAKED_AND_PENDING:      { bg: "bg-indigo-500/10", text: "text-indigo-400", dot: "bg-indigo-400" },
  PENDING_CLIENT_CONFIRMATION: { bg: "bg-yellow-500/10", text: "text-yellow-400", dot: "bg-yellow-400" },
  JOB_COMPLETED:               { bg: "bg-emerald-500/10",text: "text-emerald-400",dot: "bg-emerald-400"},
  CANCEL_REQUESTED:            { bg: "bg-orange-500/10", text: "text-orange-400", dot: "bg-orange-400" },
  DEAL_BROKEN:                 { bg: "bg-red-500/10",    text: "text-red-400",    dot: "bg-red-400"    },
  RANDOM_DISPUTED:             { bg: "bg-rose-500/10",   text: "text-rose-400",   dot: "bg-rose-400"   },
  PAYMENT_DISPUTED:            { bg: "bg-pink-500/10",   text: "text-pink-400",   dot: "bg-pink-400"   },
};

function StatusBadge({ status }: { status: string }) {
  const style = STATUS_STYLES[status] ?? STATUS_STYLES.UNALLOCATED;
  return (
    <span
      className={`inline-flex items-center shrink-0 whitespace-nowrap gap-1.5 rounded-full px-2.5 py-1 text-xs font-medium ${style.bg} ${style.text}`}
    >
      <span className={`h-1.5 w-1.5 rounded-full shrink-0 ${style.dot}`} />
      {status.replace(/_/g, " ")}
    </span>
  );
}

function truncateAddress(addr: string) {
  if (!addr || addr.length < 12) return addr || "—";
  return `${addr.slice(0, 6)}…${addr.slice(-4)}`;
}

// ── Cookie helper ───────────────────────────────────────────────────────────

function getCookie(name: string): string | undefined {
  if (typeof document === "undefined") return undefined;
  const match = document.cookie.match(new RegExp(`(?:^|; )${name}=([^;]*)`));
  return match ? decodeURIComponent(match[1]) : undefined;
}

// ── Determine staking eligibility ───────────────────────────────────────────

const CLIENT_STAKEABLE: JobStatus[] = ["AGREED", "FREELANCER_STAKED"];
const FREELANCER_STAKEABLE: JobStatus[] = ["AGREED", "CLIENT_STAKED"];

function canStake(role: string, status: JobStatus): boolean {
  if (role === "client") return CLIENT_STAKEABLE.includes(status);
  if (role === "freelancer") return FREELANCER_STAKEABLE.includes(status);
  return false;
}

// ── Staking button (isolated per-card so each card has its own tx state) ────

function StakeButton({
  job,
  role,
  onBusyChange,
}: {
  job: Job;
  role: string;
  onBusyChange?: (busy: boolean) => void;
}) {
  const queryClient = useQueryClient();
  const { data: hash, writeContract, isPending, error } = useWriteContract();
  const [stakeAmount, setStakeAmount] = useState<string>(job.payMin !== undefined && job.payMin > 0 ? job.payMin.toString() : "1");

  const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({
    hash,
  });

  // ── Auto-invalidate on success ──────────────────────────────────
  const hasInvalidated = useRef(false);

  useEffect(() => {
    if (isSuccess && !hasInvalidated.current) {
      hasInvalidated.current = true;
      queryClient.invalidateQueries({ queryKey: ["myJobs"] });
    }
  }, [isSuccess, queryClient]);

  useEffect(() => {
    hasInvalidated.current = false;
  }, [hash]);

  // ── Report busy state upstream ──────────────────────────────────
  const isBusy = isPending || isConfirming;

  useEffect(() => {
    onBusyChange?.(isBusy);
  }, [isBusy, onBusyChange]);

  const handleStake = () => {
    const functionName = role === "client" ? "addClientStake" : "addfreelancerStake";
    const val = role === "client" ? parseEther(stakeAmount || "0") : parseEther("0.01");

    writeContract({
      abi: FREELANCE_ESCROW_ABI,
      address: job.contractAddress as `0x${string}`,
      functionName,
      value: val,
    });
  };

  if (isSuccess) {
    return (
      <div className="mt-4 flex items-center gap-2 rounded-xl border border-emerald-500/20 bg-emerald-500/10 px-4 py-2.5">
        <span className="h-2 w-2 rounded-full bg-emerald-400" />
        <span className="text-xs font-medium text-emerald-400">
          Stake confirmed!
        </span>
      </div>
    );
  }

  const isWaiting = isPending || isConfirming;

  return (
    <div className="mt-4 flex flex-col gap-3">
      {role === "client" && (
        <div className="flex items-center justify-between gap-3 bg-slate-900/60 rounded-xl px-4 py-2 border border-slate-700/60">
          <label className="text-xs font-semibold text-slate-400">Stake (ETH):</label>
          <input 
            type="number" 
            step="0.01"
            min="0"
            value={stakeAmount}
            onChange={(e) => setStakeAmount(e.target.value)}
            disabled={isWaiting}
            className="w-24 bg-transparent text-sm text-right text-slate-200 outline-none placeholder-slate-600 font-mono disabled:opacity-50"
          />
        </div>
      )}
      <button
        onClick={handleStake}
        disabled={isWaiting}
        className="flex w-full items-center justify-center gap-2 rounded-xl px-4 py-2.5 text-sm font-semibold text-white transition-all duration-200 disabled:cursor-not-allowed disabled:opacity-40"
        style={{
          background: `rgba(var(--vault-accent), 0.8)`,
          boxShadow: `0 0 20px rgba(var(--vault-accent), 0.2)`,
        }}
        onMouseEnter={(e) => {
          if (!isWaiting) {
            e.currentTarget.style.background = `rgba(var(--vault-accent), 0.9)`;
            e.currentTarget.style.boxShadow = `0 0 28px rgba(var(--vault-accent), 0.3)`;
          }
        }}
        onMouseLeave={(e) => {
          e.currentTarget.style.background = `rgba(var(--vault-accent), 0.8)`;
          e.currentTarget.style.boxShadow = `0 0 20px rgba(var(--vault-accent), 0.2)`;
        }}
      >
        {isWaiting ? (
          <>
            <Loader2 className="h-4 w-4 animate-spin" />
            {isPending ? "Confirming in Wallet…" : "Waiting for chain…"}
          </>
        ) : (
          <>
            <Wallet className="h-4 w-4" />
            Stake Funds ({role === "client" ? stakeAmount || "0" : "0.01"} ETH)
          </>
        )}
      </button>
      {error && (
        <p className="mt-2 text-xs text-red-400/80 truncate" title={error.message}>
          {error.message.length > 80
            ? error.message.slice(0, 80) + "…"
            : error.message}
        </p>
      )}
    </div>
  );
}

// ── Create Job Modal (Client only) ──────────────────────────────────────────

function CreateJobModal({
  isOpen,
  onClose,
}: {
  isOpen: boolean;
  onClose: () => void;
}) {
  const { address } = useAccount();
  const queryClient = useQueryClient();
  const [isCreating, setIsCreating] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState(false);

  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [deadline, setDeadline] = useState("");
  const [contactEmail, setContactEmail] = useState("");
  const [payMin, setPayMin] = useState<number | "">("");
  const [payMax, setPayMax] = useState<number | "">("");

  const resetForm = () => {
    setTitle("");
    setDescription("");
    setDeadline("");
    setContactEmail("");
    setPayMin("");
    setPayMax("");
    setError(null);
    setSuccess(false);
  };

  const handleClose = () => {
    resetForm();
    onClose();
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!address) {
      setError("Wallet not connected.");
      return;
    }

    if (!title.trim()) {
      setError("Title is required.");
      return;
    }

    setIsCreating(true);
    setError(null);
    setSuccess(false);

    try {
      const jobId = new ObjectId().toHexString();

      await axios.post(
        `${process.env.NEXT_PUBLIC_API_URL}/api/post/job/create`,
        {
          clientEthAccount: address,
          jobId,
          title: title.trim(),
          description: description.trim(),
          deadline: deadline.trim(),
          contactEmail: contactEmail.trim(),
          payMin: Number(payMin),
          payMax: Number(payMax),
        },
        { withCredentials: true }
      );

      setSuccess(true);
      await queryClient.invalidateQueries({ queryKey: ["myJobs"] });

      // Auto-close after brief success message
      setTimeout(() => {
        handleClose();
      }, 1500);
    } catch (err) {
      setError(extractErrorMsg(err, "Unknown error"));
    } finally {
      setIsCreating(false);
    }
  };

  if (!isOpen) return null;

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center p-4"
      style={{ backdropFilter: "blur(8px)" }}
    >
      {/* Backdrop */}
      <div
        className="absolute inset-0 bg-black/60"
        onClick={handleClose}
      />

      {/* Modal */}
      <div
        className="relative w-full max-w-3xl rounded-2xl p-6 shadow-2xl"
        style={{
          background: `linear-gradient(145deg, rgba(15, 23, 42, 0.95), rgba(10, 14, 26, 0.98))`,
          border: `1px solid rgba(var(--vault-accent), 0.20)`,
          boxShadow: `0 0 60px rgba(var(--vault-accent), 0.10), 0 25px 50px rgba(0,0,0,0.5)`,
        }}
      >
        {/* Close button */}
        <button
          onClick={handleClose}
          className="absolute top-4 right-4 rounded-lg p-1.5 text-slate-500 transition-colors hover:bg-slate-800 hover:text-slate-300"
        >
          <X className="h-4 w-4" />
        </button>

        {/* Header */}
        <div className="mb-6">
          <div className="flex items-center gap-3 mb-1">
            <div
              className="flex h-9 w-9 items-center justify-center rounded-xl"
              style={{
                background: `rgba(var(--vault-accent), 0.12)`,
                border: `1px solid rgba(var(--vault-accent), 0.20)`,
              }}
            >
              <Plus className="h-4 w-4" style={{ color: `rgba(var(--vault-accent), 1)` }} />
            </div>
            <h2 className="text-lg font-bold tracking-tight text-slate-100">
              Create New Job
            </h2>
          </div>
          <p className="text-xs text-slate-500 pl-12">
            Post a new escrow agreement to the marketplace
          </p>
        </div>

        {/* Form */}
        <form onSubmit={handleSubmit} className="space-y-4">
          {/* Title */}
          <div>
            <label className="mb-1.5 flex items-center gap-1.5 text-xs font-medium text-slate-400">
              <Type className="h-3 w-3" />
              Title <span className="text-red-400">*</span>
            </label>
            <input
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              placeholder="e.g. Full-Stack dApp Development"
              className="w-full rounded-xl border border-slate-700/60 bg-slate-900/60 px-4 py-2.5 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
              required
            />
          </div>

          {/* Description */}
          <div>
            <label className="mb-1.5 flex items-center gap-1.5 text-xs font-medium text-slate-400">
              <FileText className="h-3 w-3" />
              Description
            </label>
            <textarea
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              placeholder="Describe the scope of work, deliverables, and milestones…"
              rows={3}
              className="w-full resize-y rounded-xl border border-slate-700/60 bg-slate-900/60 px-4 py-2.5 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
            />
          </div>

          {/* Deadline */}
          <div>
            <label className="mb-1.5 flex items-center gap-1.5 text-xs font-medium text-slate-400">
              <Calendar className="h-3 w-3" />
              Deadline
            </label>
            <input
              type="date"
              value={deadline}
              onChange={(e) => setDeadline(e.target.value)}
              className="w-full rounded-xl border border-slate-700/60 bg-slate-900/60 px-4 py-2.5 text-sm text-slate-200 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)] [color-scheme:dark]"
            />
          </div>

          {/* Pay Range (Min/Max ETH) */}
          <div className="grid grid-cols-2 gap-3">
            <div>
              <label className="mb-1.5 flex items-center gap-1.5 text-xs font-medium text-slate-400">
                <DollarSign className="h-3 w-3" />
                Min ETH
              </label>
              <input
                type="number"
                step="0.01"
                min="0"
                value={payMin}
                onChange={(e) => setPayMin(e.target.value ? Number(e.target.value) : "")}
                placeholder="e.g. 1.5"
                className="w-full rounded-xl border border-slate-700/60 bg-slate-900/60 px-4 py-2.5 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
              />
            </div>
            <div>
              <label className="mb-1.5 flex items-center gap-1.5 text-xs font-medium text-slate-400">
                <DollarSign className="h-3 w-3" />
                Max ETH
              </label>
              <input
                type="number"
                step="0.01"
                min="0"
                value={payMax}
                onChange={(e) => setPayMax(e.target.value ? Number(e.target.value) : "")}
                placeholder="e.g. 3.0"
                className="w-full rounded-xl border border-slate-700/60 bg-slate-900/60 px-4 py-2.5 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
              />
            </div>
          </div>

          {/* Contact Email */}
          <div>
            <label className="mb-1.5 flex items-center gap-1.5 text-xs font-medium text-slate-400">
              <Mail className="h-3 w-3" />
              Contact Email
            </label>
            <input
              type="email"
              value={contactEmail}
              onChange={(e) => setContactEmail(e.target.value)}
              placeholder="you@example.com"
              className="w-full rounded-xl border border-slate-700/60 bg-slate-900/60 px-4 py-2.5 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-[rgba(var(--vault-accent),0.5)] focus:ring-1 focus:ring-[rgba(var(--vault-accent),0.25)]"
            />
          </div>

          {/* Error / Success */}
          {error && (
            <div className="rounded-xl border border-red-500/20 bg-red-500/10 px-4 py-2.5">
              <p className="text-xs text-red-400">{error}</p>
            </div>
          )}
          {success && (
            <div className="flex items-center gap-2 rounded-xl border border-emerald-500/20 bg-emerald-500/10 px-4 py-2.5">
              <span className="h-2 w-2 rounded-full bg-emerald-400" />
              <span className="text-xs font-medium text-emerald-400">
                Job created successfully!
              </span>
            </div>
          )}

          {/* Submit */}
          <button
            type="submit"
            disabled={isCreating}
            className="flex w-full items-center justify-center gap-2 rounded-xl px-4 py-3 text-sm font-semibold text-white transition-all duration-200 disabled:cursor-not-allowed disabled:opacity-50"
            style={{
              background: `rgba(var(--vault-accent), 0.85)`,
              boxShadow: `0 0 24px rgba(var(--vault-accent), 0.20)`,
            }}
            onMouseEnter={(e) => {
              if (!isCreating) {
                e.currentTarget.style.background = `rgba(var(--vault-accent), 1)`;
                e.currentTarget.style.boxShadow = `0 0 32px rgba(var(--vault-accent), 0.35)`;
                e.currentTarget.style.transform = "translateY(-1px)";
              }
            }}
            onMouseLeave={(e) => {
              e.currentTarget.style.background = `rgba(var(--vault-accent), 0.85)`;
              e.currentTarget.style.boxShadow = `0 0 24px rgba(var(--vault-accent), 0.20)`;
              e.currentTarget.style.transform = "translateY(0)";
            }}
          >
            {isCreating ? (
              <>
                <Loader2 className="h-4 w-4 animate-spin" />
                Creating Job…
              </>
            ) : (
              <>
                <Plus className="h-4 w-4" />
                Create Job
              </>
            )}
          </button>
        </form>
      </div>
    </div>
  );
}

// ── Dashboard page ──────────────────────────────────────────────────────────

export default function DashboardPage() {
  const [role, setRole] = useState<string>("");
  const [activeTab, setActiveTab] = useState<"jobs" | "profile">("jobs");
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [busyCards, setBusyCards] = useState<Record<string, boolean>>({});
  const [reviewJob, setReviewJob] = useState<Job | null>(null);
  const [isReviewModalOpen, setIsReviewModalOpen] = useState(false);

  // Callback factory for per-card, per-component busy tracking
  const makeCardBusyCb = useCallback(
    (key: string) => (busy: boolean) => {
      setBusyCards((prev) => {
        if (prev[key] === busy) return prev;
        return { ...prev, [key]: busy };
      });
    },
    []
  );

  const isCardBusy = (jobId: string) => 
    Object.keys(busyCards).some(k => k.startsWith(jobId) && busyCards[k]);

  useEffect(() => {
    setRole(getCookie("role") ?? "");
  }, []);

  const { data: jobs, isLoading, isError, isRefetching } = useQuery<Job[]>({
    queryKey: ["myJobs"],
    queryFn: async () => {
      const res = await axios.get(
        `${process.env.NEXT_PUBLIC_API_URL}/api/get/jobs/me`,
        { withCredentials: true }
      );
      return res.data;
    },
  });

  // ── Deduplication: filter by unique job.id to prevent React key collisions ──
  const uniqueJobs = jobs
    ? jobs.filter((v, i, a) => a.findIndex((t) => t.id === v.id) === i)
    : [];

  return (
    <div
      className="min-h-screen font-[family-name:var(--font-geist-sans)]"
      style={{
        background: `linear-gradient(to bottom, var(--vault-bg-start), var(--vault-bg-mid), var(--vault-bg-end))`,
      }}
    >
      <AppNavbar />

      {/* Subtle grid overlay */}
      <div
        className="pointer-events-none fixed inset-0"
        style={{
          backgroundImage:
            `linear-gradient(var(--vault-grid) 1px, transparent 1px), linear-gradient(90deg, var(--vault-grid) 1px, transparent 1px)`,
          backgroundSize: "64px 64px",
        }}
      />

      <main className="relative z-10 mx-auto max-w-6xl px-6 py-10">
        {/* Header */}
        <div className="mb-8">
          <div className="flex items-center gap-3 mb-2">
            <div
              className="flex h-10 w-10 items-center justify-center rounded-xl"
              style={{
                background: `rgba(var(--vault-accent), 0.10)`,
                border: `1px solid rgba(var(--vault-accent), 0.20)`,
              }}
            >
              <Briefcase className="h-5 w-5" style={{ color: `rgba(var(--vault-accent), 1)` }} />
            </div>
            <h1 className="text-2xl font-bold tracking-tight text-slate-100">
              My Jobs
            </h1>

            {/* Sync spinner — visible when JIT background refetch is active */}
            {isRefetching && (
              <div className="ml-2 flex items-center gap-1.5 rounded-full border border-slate-700/50 bg-slate-800/60 px-3 py-1">
                <Loader2
                  className="h-3.5 w-3.5 animate-spin"
                  style={{ color: `rgba(var(--vault-accent), 0.8)` }}
                />
                <span className="text-[10px] font-medium text-slate-500">
                  Syncing…
                </span>
              </div>
            )}

            {role && (
              <span className="ml-auto inline-flex items-center gap-1.5 rounded-full border border-slate-700/50 bg-slate-800/60 px-3 py-1 text-xs font-medium text-slate-400">
                <span
                  className="h-1.5 w-1.5 rounded-full"
                  style={{ background: `rgba(var(--vault-accent), 1)` }}
                />
                {role.charAt(0).toUpperCase() + role.slice(1)}
              </span>
            )}
          </div>
          <p className="text-sm text-slate-500 pl-[52px]">
            All escrow contracts associated with your account
          </p>
        </div>

        {/* Tab Navigation for Freelancers */}
        {role === "freelancer" && (
          <div className="mb-8 flex gap-2 border-b border-slate-800/60 pb-px">
            <button
              onClick={() => setActiveTab("jobs")}
              className={`px-4 py-2 text-sm font-medium transition-colors ${
                activeTab === "jobs"
                  ? "border-b-2 border-[rgba(var(--vault-accent),1)] text-slate-100"
                  : "border-b-2 border-transparent text-slate-400 hover:text-slate-300"
              }`}
            >
              My Jobs
            </button>
            <button
              onClick={() => setActiveTab("profile")}
              className={`px-4 py-2 text-sm font-medium transition-colors ${
                activeTab === "profile"
                  ? "border-b-2 border-[rgba(var(--vault-accent),1)] text-slate-100"
                  : "border-b-2 border-transparent text-slate-400 hover:text-slate-300"
              }`}
            >
              My Profile
            </button>
          </div>
        )}

        {activeTab === "profile" && role === "freelancer" ? (
          <FreelancerProfile />
        ) : (
          <>
            {/* Create Job — Client only */}
            {role === "client" && (
              <div className="mb-8">
            <button
              onClick={() => setIsModalOpen(true)}
              className="flex items-center gap-2 rounded-xl px-5 py-2.5 text-sm font-semibold text-white transition-all duration-200"
              style={{
                background: `rgba(var(--vault-accent), 0.85)`,
                boxShadow: `0 0 24px rgba(var(--vault-accent), 0.20)`,
              }}
              onMouseEnter={(e) => {
                e.currentTarget.style.background = `rgba(var(--vault-accent), 1)`;
                e.currentTarget.style.boxShadow = `0 0 32px rgba(var(--vault-accent), 0.35)`;
                e.currentTarget.style.transform = "translateY(-1px)";
              }}
              onMouseLeave={(e) => {
                e.currentTarget.style.background = `rgba(var(--vault-accent), 0.85)`;
                e.currentTarget.style.boxShadow = `0 0 24px rgba(var(--vault-accent), 0.20)`;
                e.currentTarget.style.transform = "translateY(0)";
              }}
            >
              <Plus className="h-4 w-4" />
              Create New Job
            </button>
          </div>
        )}

        {/* Create Job Modal */}
        <CreateJobModal
          isOpen={isModalOpen}
          onClose={() => setIsModalOpen(false)}
        />

        {/* Loading */}
        {isLoading && (
          <div className="flex flex-col items-center justify-center gap-3 py-32">
            <Loader2 className="h-8 w-8 animate-spin" style={{ color: `rgba(var(--vault-accent), 0.6)` }} />
            <span className="text-sm text-slate-500">Loading your jobs…</span>
          </div>
        )}

        {/* Error */}
        {isError && (
          <div className="flex flex-col items-center justify-center gap-3 py-32">
            <div className="flex h-14 w-14 items-center justify-center rounded-full bg-red-500/10 border border-red-500/20">
              <span className="text-2xl">⚠</span>
            </div>
            <p className="text-sm text-slate-400">
              Failed to load jobs. Make sure you are logged in.
            </p>
          </div>
        )}

        {/* Empty */}
        {!isLoading && !isError && uniqueJobs.length === 0 && (
          <div className="flex flex-col items-center justify-center gap-4 py-32">
            <div className="flex h-16 w-16 items-center justify-center rounded-2xl bg-slate-800/60 border border-slate-700/40">
              <Inbox className="h-7 w-7 text-slate-600" />
            </div>
            <div className="text-center">
              <p className="text-sm font-medium text-slate-400">No jobs yet</p>
              <p className="mt-1 text-xs text-slate-600">
                {role === "client"
                  ? "Click \"Create New Job\" above to post your first escrow agreement."
                  : "Accept a job from the marketplace to get started."}
              </p>
            </div>
          </div>
        )}

        {/* Job Cards */}
        {!isLoading && !isError && uniqueJobs.length > 0 && (
          <div className="grid gap-4 sm:grid-cols-2">
            {uniqueJobs.map((job) => (
              <div
                key={job.id}
                className="group relative rounded-2xl backdrop-blur-sm p-5 transition-all duration-300"
                style={{
                  border: `1px solid rgba(var(--vault-accent), 0.10)`,
                  background: `rgba(15, 23, 42, 0.70)`,
                }}
                onMouseEnter={(e) => {
                  e.currentTarget.style.borderColor = `rgba(var(--vault-accent), 0.25)`;
                  e.currentTarget.style.boxShadow = `0 0 40px var(--vault-card-hover)`;
                }}
                onMouseLeave={(e) => {
                  e.currentTarget.style.borderColor = `rgba(var(--vault-accent), 0.10)`;
                  e.currentTarget.style.boxShadow = "none";
                }}
              >
                {/* ── Per-card action spinner (top-right) ─────────────── */}
                {(isCardBusy(job.id) || isRefetching) && (
                  <div className="absolute top-3 right-3 z-10 flex items-center gap-1.5 rounded-full border border-slate-700/50 bg-slate-900/80 px-2.5 py-1 backdrop-blur-sm">
                    <Loader2
                      className="h-3 w-3 animate-spin"
                      style={{ color: `rgba(var(--vault-accent), 0.9)` }}
                    />
                    <span className="text-[9px] font-medium text-slate-500">
                      {isCardBusy(job.id) ? "Processing…" : "Syncing…"}
                    </span>
                  </div>
                )}

                {/* Title + Status Row */}
                <div className="mb-3 flex items-start justify-between gap-2">
                  <h3 className="text-base font-semibold text-slate-100 leading-snug">
                    {job.title || "Untitled Job"}
                  </h3>
                  <StatusBadge status={job.status} />
                </div>

                {/* Description */}
                {job.description && (
                  <p className="mb-3 text-xs text-slate-400 leading-relaxed line-clamp-2">
                    {job.description}
                  </p>
                )}

                {/* Pay Range & Deadline pills */}
                <div className="mb-4 flex flex-wrap gap-2">
                  {(job.payMin !== undefined || job.payMax !== undefined) && (
                    <span className="inline-flex items-center gap-1 rounded-lg bg-emerald-500/10 px-2 py-0.5 text-[10px] font-medium text-emerald-400">
                      <DollarSign className="h-2.5 w-2.5" />
                      {job.payMin === job.payMax ? `${job.payMin} ETH` : `${job.payMin} - ${job.payMax} ETH`}
                    </span>
                  )}
                  {job.deadline && (
                    <span className="inline-flex items-center gap-1 rounded-lg bg-blue-500/10 px-2 py-0.5 text-[10px] font-medium text-blue-400">
                      <Calendar className="h-2.5 w-2.5" />
                      {job.deadline}
                    </span>
                  )}
                </div>

                {/* Contract Address */}
                <div className="mb-4">
                  <span className="text-[10px] font-medium uppercase tracking-widest text-slate-600">
                    Contract
                  </span>
                  <div className="mt-1 flex items-center gap-2">
                    <code className="font-mono text-sm text-slate-300">
                      {truncateAddress(job.contractAddress)}
                    </code>
                    {job.contractAddress && (
                      <ExternalLink className="h-3 w-3 text-slate-600 opacity-0 transition-opacity group-hover:opacity-100" />
                    )}
                  </div>
                </div>

                {/* IDs */}
                <div className="grid grid-cols-2 gap-3 border-t border-slate-800/60 pt-4">
                  <div>
                    <span className="text-[10px] font-medium uppercase tracking-widest text-slate-600">
                      Client
                    </span>
                    <p className="mt-0.5 font-mono text-xs text-slate-400">
                      {truncateAddress(job.clientId)}
                    </p>
                  </div>
                  <div>
                    <span className="text-[10px] font-medium uppercase tracking-widest text-slate-600">
                      Freelancer
                    </span>
                    <p className="mt-0.5 font-mono text-xs text-slate-400">
                      {truncateAddress(job.freelancerId)}
                    </p>
                  </div>
                </div>

                {/* View Applicants Button - only for unallocated jobs for client */}
                {role === "client" && job.status === "UNALLOCATED" && job.applicants && job.applicants.length > 0 && (
                  <div className="mt-4">
                    <button
                      onClick={() => {
                        setReviewJob(job);
                        setIsReviewModalOpen(true);
                      }}
                      className="flex w-full items-center justify-center gap-2 rounded-xl px-4 py-2.5 text-sm font-semibold text-white transition-all duration-200 hover:bg-[rgba(var(--vault-accent),0.9)]"
                      style={{
                        background: `rgba(var(--vault-accent), 0.8)`,
                        boxShadow: `0 0 20px rgba(var(--vault-accent), 0.2)`,
                      }}
                    >
                      View Applicants ({job.applicants.length})
                    </button>
                  </div>
                )}

                {/* Stake Button — conditional on role + status */}
                {role && canStake(role, job.status) && (
                  <StakeButton
                    job={job}
                    role={role}
                    onBusyChange={makeCardBusyCb(`${job.id}-stake`)}
                  />
                )}

                {/* ── Status-dependent action buttons ── */}
                <div className="mt-4 border-t border-slate-800/60 pt-4">
                  <JobCardActions
                    job={job}
                    role={role}
                    onBusyChange={makeCardBusyCb(`${job.id}-actions`)}
                  />
                </div>

                {/* Glow accent on hover */}
                <div
                  className="pointer-events-none absolute inset-0 rounded-2xl opacity-0 transition-opacity duration-300 group-hover:opacity-100"
                  style={{
                    background: `linear-gradient(to bottom, rgba(var(--vault-accent), 0.02), transparent)`,
                  }}
                />
              </div>
            ))}
          </div>
        )}
          </>
        )}
        <ApplicantReviewModal
          isOpen={isReviewModalOpen}
          onClose={() => {
            setIsReviewModalOpen(false);
            setReviewJob(null);
          }}
          job={reviewJob}
        />
      </main>
    </div>
  );
}