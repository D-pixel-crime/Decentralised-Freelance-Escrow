"use client";

import React, { useCallback, useEffect, useState } from "react";
import type { Job } from "@/types/job";
import DeliveryActions from "./DeliveryActions";
import DangerZoneActions from "./DangerZoneActions";

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

  // Terminal / irrelevant states — nothing to render
  const terminalStatuses = ["JOB_COMPLETED", "DEAL_BROKEN", "PAYMENT_DISPUTED"];
  if (terminalStatuses.includes(job.status) && job.status !== "PAYMENT_DISPUTED") {
    // PAYMENT_DISPUTED might still show danger zone info
  }

  return (
    <>
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
