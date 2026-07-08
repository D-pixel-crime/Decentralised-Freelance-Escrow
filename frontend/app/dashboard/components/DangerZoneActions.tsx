"use client";

import React, { useCallback, useEffect, useState } from "react";
import {
  Loader2,
  ChevronDown,
  ChevronUp,
  Ban,
  Undo2,
  AlertTriangle,
} from "lucide-react";
import { useEscrowAction } from "./useEscrowAction";
import type { Job, JobStatus } from "@/types/job";

// ── Statuses where Break Deal is available ──────────────────────────────────

const BREAKABLE_STATUSES: JobStatus[] = [
  "AGREED",
  "CLIENT_STAKED",
  "FREELANCER_STAKED",
  "ALL_STAKED_AND_PENDING",
  "PENDING_CLIENT_CONFIRMATION",
];

// ── Statuses where Raise Dispute is available ───────────────────────────────

const DISPUTABLE_STATUSES: JobStatus[] = [
  "AGREED",
  "CLIENT_STAKED",
  "FREELANCER_STAKED",
  "ALL_STAKED_AND_PENDING",
  "PENDING_CLIENT_CONFIRMATION",
];

// ── Small danger button ─────────────────────────────────────────────────────

function DangerButton({
  onClick,
  disabled,
  isPending,
  isConfirming,
  isSuccess,
  icon: Icon,
  label,
  successLabel,
  color = "red",
}: {
  onClick: () => void;
  disabled: boolean;
  isPending: boolean;
  isConfirming: boolean;
  isSuccess: boolean;
  icon: React.ComponentType<{ className?: string }>;
  label: string;
  successLabel: string;
  color?: "red" | "orange" | "amber";
}) {
  const isWaiting = isPending || isConfirming;

  const colorMap = {
    red: {
      border: "border-red-500/20",
      text: "text-red-400",
      bg: "bg-red-500/8",
      hoverBg: "hover:bg-red-500/15",
      dot: "bg-red-400",
    },
    orange: {
      border: "border-orange-500/20",
      text: "text-orange-400",
      bg: "bg-orange-500/8",
      hoverBg: "hover:bg-orange-500/15",
      dot: "bg-orange-400",
    },
    amber: {
      border: "border-amber-500/20",
      text: "text-amber-400",
      bg: "bg-amber-500/8",
      hoverBg: "hover:bg-amber-500/15",
      dot: "bg-amber-400",
    },
  };
  const c = colorMap[color];

  if (isSuccess) {
    return (
      <div className={`flex items-center gap-2 rounded-lg border ${c.border} ${c.bg} px-3 py-1.5`}>
        <span className={`h-1.5 w-1.5 rounded-full ${c.dot}`} />
        <span className={`text-[11px] font-medium ${c.text}`}>{successLabel}</span>
      </div>
    );
  }

  return (
    <button
      onClick={onClick}
      disabled={disabled || isWaiting}
      className={`flex items-center gap-1.5 rounded-lg border ${c.border} ${c.bg} ${c.hoverBg} px-3 py-1.5 text-[11px] font-medium ${c.text} transition-all duration-200 disabled:cursor-not-allowed disabled:opacity-40`}
    >
      {isWaiting ? (
        <>
          <Loader2 className="h-3 w-3 animate-spin" />
          {isPending ? "Wallet…" : "Chain…"}
        </>
      ) : (
        <>
          <Icon className="h-3 w-3" />
          {label}
        </>
      )}
    </button>
  );
}

// ── DangerZoneActions ───────────────────────────────────────────────────────

export default function DangerZoneActions({
  job,
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  role,
  onBusyChange,
}: {
  job: Job;
  role: string;
  onBusyChange: (busy: boolean) => void;
}) {
  const [isOpen, setIsOpen] = useState(false);
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

  useEffect(() => {
    const anyBusy = Object.values(busyMap).some(Boolean);
    onBusyChange(anyBusy);
  }, [busyMap, onBusyChange]);

  const breakDeal = useEscrowAction(
    job.contractAddress,
    "breakDeal",
    makeBusyCb("breakDeal")
  );

  const cancelDealBreak = useEscrowAction(
    job.contractAddress,
    "cancelDealBreak",
    makeBusyCb("cancelDealBreak")
  );

  const raiseDispute = useEscrowAction(
    job.contractAddress,
    "raiseDispute",
    makeBusyCb("raiseDispute")
  );

  const anyBusy = Object.values(busyMap).some(Boolean);

  // ── Determine which actions to show ──────────────────────────────
  const showBreakDeal = BREAKABLE_STATUSES.includes(job.status);
  const showCancelActions = job.status === "CANCEL_REQUESTED";
  const showDispute =
    DISPUTABLE_STATUSES.includes(job.status) &&
    job.status !== "RANDOM_DISPUTED" &&
    job.status !== "PAYMENT_DISPUTED";

  // Terminal states — no danger zone needed
  if (!showBreakDeal && !showCancelActions && !showDispute) {
    return null;
  }

  const aggregatedError = breakDeal.error || cancelDealBreak.error || raiseDispute.error;

  return (
    <div className="mt-4 border-t border-slate-800/40 pt-3">
      {/* Toggle header */}
      <button
        onClick={() => setIsOpen((prev) => !prev)}
        className="flex w-full items-center gap-1.5 text-[10px] font-medium uppercase tracking-widest text-slate-600 transition-colors hover:text-slate-400"
      >
        <AlertTriangle className="h-3 w-3" />
        Danger Zone
        {isOpen ? (
          <ChevronUp className="ml-auto h-3 w-3" />
        ) : (
          <ChevronDown className="ml-auto h-3 w-3" />
        )}
      </button>

      {/* Collapsible body */}
      {isOpen && (
        <div className="mt-2.5 flex flex-wrap gap-2">
          {/* Break Deal — available on active statuses */}
          {showBreakDeal && (
            <DangerButton
              onClick={breakDeal.execute}
              disabled={anyBusy}
              isPending={breakDeal.isPending}
              isConfirming={breakDeal.isConfirming}
              isSuccess={breakDeal.isSuccess}
              icon={Ban}
              label="Break Deal"
              successLabel="Cancel requested"
              color="orange"
            />
          )}

          {/* Cancel Requested state — Break Deal (confirm) + Revert Cancel */}
          {showCancelActions && (
            <>
              <DangerButton
                onClick={breakDeal.execute}
                disabled={anyBusy}
                isPending={breakDeal.isPending}
                isConfirming={breakDeal.isConfirming}
                isSuccess={breakDeal.isSuccess}
                icon={Ban}
                label="Confirm Break"
                successLabel="Deal broken"
                color="red"
              />
              <DangerButton
                onClick={cancelDealBreak.execute}
                disabled={anyBusy}
                isPending={cancelDealBreak.isPending}
                isConfirming={cancelDealBreak.isConfirming}
                isSuccess={cancelDealBreak.isSuccess}
                icon={Undo2}
                label="Revert Cancel"
                successLabel="Cancel reverted"
                color="amber"
              />
            </>
          )}

          {/* Raise Dispute — available on active non-disputed states */}
          {showDispute && (
            <DangerButton
              onClick={raiseDispute.execute}
              disabled={anyBusy}
              isPending={raiseDispute.isPending}
              isConfirming={raiseDispute.isConfirming}
              isSuccess={raiseDispute.isSuccess}
              icon={AlertTriangle}
              label="Raise Dispute"
              successLabel="Dispute raised"
              color="red"
            />
          )}

          {/* Error line */}
          {aggregatedError && (
            <p className="w-full text-xs text-red-400/80 truncate mt-1" title={aggregatedError.message}>
              {aggregatedError.message.length > 80
                ? aggregatedError.message.slice(0, 80) + "…"
                : aggregatedError.message}
            </p>
          )}
        </div>
      )}
    </div>
  );
}
