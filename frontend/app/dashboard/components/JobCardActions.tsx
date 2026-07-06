"use client";

import React, { useCallback, useEffect, useState } from "react";
import type { Job } from "@/types/job";
import DeliveryActions from "./DeliveryActions";
import DangerZoneActions from "./DangerZoneActions";
import SettlementPanel from "./SettlementPanel";
import { User, Gavel } from "lucide-react";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";

/**
 * JobCardActions — orchestrator that renders the correct action groups
 * based on job.status and user role, and aggregates busy state upward.
 */
export default function JobCardActions({
  job,
  role,
  onBusyChange,
}: {
  job: Job;
  role: string;
  onBusyChange: (busy: boolean) => void;
}) {
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

  // Bubble aggregated busy state to the card
  useEffect(() => {
    const anyBusy = Object.values(busyMap).some(Boolean);
    onBusyChange(anyBusy);
  }, [busyMap, onBusyChange]);

  if (role === "arbitrator") {
    if (job.status === "RANDOM_DISPUTED" || job.status === "PAYMENT_DISPUTED") {
      return <SettlementPanel job={job} onBusyChange={makeBusyCb("settlement")} />;
    }
    if (job.status === "DEAL_BROKEN") {
      return (
        <div className="mt-4 rounded-xl border border-slate-700/60 bg-slate-800/40 px-4 py-3 text-center">
          <span className="text-sm text-slate-400">This deal has been broken.</span>
        </div>
      );
    }

    if (job.status === "JOB_COMPLETED") {
      return (
        <div className="mt-4 rounded-xl border border-slate-700/60 bg-slate-800/40 px-4 py-3 text-center">
          <span className="text-sm text-slate-400">This job has been completed.</span>
        </div>
      );
    }

    return (
      <div className="mt-4 rounded-xl border border-slate-700/60 bg-slate-800/40 px-4 py-3 text-center">
        <span className="text-sm text-slate-400">Waiting for parties to reach {job.status === "UNALLOCATED" ? "AGREED" : "a disputed"} state...</span>
      </div>
    );
  }

  // Terminal / irrelevant states — nothing to render
  const terminalStatuses = ["JOB_COMPLETED", "DEAL_BROKEN", "PAYMENT_DISPUTED"];
  if (terminalStatuses.includes(job.status) && job.status !== "PAYMENT_DISPUTED") {
    // PAYMENT_DISPUTED might still show danger zone info
  }

  // ── Dispute Info for Client & Freelancer ──
  const isDisputed = job.status === "RANDOM_DISPUTED" || job.status === "PAYMENT_DISPUTED";
  
  const { data: arbitratorContact } = useQuery({
    queryKey: ["contact", job.arbitratorEth, "arbitrator"],
    queryFn: async () => {
      const res = await axios.get(
        `${process.env.NEXT_PUBLIC_API_URL}/api/get/contact/${job.arbitratorEth}?role=arbitrator`,
        { withCredentials: true }
      );
      return res.data;
    },
    enabled: isDisputed && !!job.arbitratorEth && role !== "arbitrator",
  });

  return (
    <>
      {isDisputed && role !== "arbitrator" && (
        <div className="mt-4 rounded-xl border border-amber-500/20 bg-amber-500/5 p-4">
          <h4 className="mb-3 text-xs font-semibold text-amber-400 flex items-center gap-2 uppercase tracking-wider">
            <Gavel className="h-4 w-4" />
            Arbitrator Assigned
          </h4>
          <div className="space-y-3">
            <div className="flex flex-col gap-1 text-sm text-slate-300">
              <span className="text-xs text-slate-500">Arbitrator Address:</span>
              <span className="font-mono text-xs">{job.arbitratorEth}</span>
            </div>
            <div className="flex items-center gap-2 text-sm text-slate-300">
              <User className="h-4 w-4 text-slate-500" />
              <span className="text-slate-500">Contact:</span> {arbitratorContact?.email || <span className="text-slate-600 italic">No email provided</span>}
            </div>
          </div>
        </div>
      )}

      {/* Delivery actions: Request Payment / Approve / Reject */}
      <DeliveryActions
        job={job}
        role={role}
        onBusyChange={makeBusyCb("delivery")}
      />

      {/* Danger zone: Break Deal / Revert Cancel / Raise Dispute */}
      <DangerZoneActions
        job={job}
        role={role}
        onBusyChange={makeBusyCb("danger")}
      />
    </>
  );
}
