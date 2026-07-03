"use client";

import React, { useState } from "react";
import { useReadContract, useWriteContract, useWaitForTransactionReceipt } from "wagmi";
import { parseEther, formatEther } from "viem";
import { FREELANCE_ESCROW_ABI } from "@/constants/contract";
import { Loader2, ShieldCheck, DollarSign } from "lucide-react";
import type { Job } from "@/types/job";
import { useQueryClient } from "@tanstack/react-query";

interface SettlementPanelProps {
  job: Job;
  onBusyChange: (busy: boolean) => void;
}

export default function SettlementPanel({ job, onBusyChange }: SettlementPanelProps) {
  const [clientRefund, setClientRefund] = useState("");
  const [freelancerPayment, setFreelancerPayment] = useState("");

  const queryClient = useQueryClient();

  // Fetch balances
  const { data: freelancerStake } = useReadContract({
    address: job.contractAddress as `0x${string}`,
    abi: FREELANCE_ESCROW_ABI,
    functionName: "getFreelancerStake",
  });

  const { data: clientStake } = useReadContract({
    address: job.contractAddress as `0x${string}`,
    abi: FREELANCE_ESCROW_ABI,
    functionName: "getClientStake",
  });

  const { data: hash, writeContract, isPending, error } = useWriteContract();
  const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({ hash });

  const totalBalance = (freelancerStake ? BigInt(freelancerStake.toString()) : BigInt(0)) + 
                       (clientStake ? BigInt(clientStake.toString()) : BigInt(0));

  const handleSettle = () => {
    if (!clientRefund || !freelancerPayment) return;
    
    writeContract({
      address: job.contractAddress as `0x${string}`,
      abi: FREELANCE_ESCROW_ABI,
      functionName: "resolveDispute",
      args: [parseEther(clientRefund), parseEther(freelancerPayment)],
    });
  };

  React.useEffect(() => {
    onBusyChange(isPending || isConfirming);
  }, [isPending, isConfirming, onBusyChange]);

  React.useEffect(() => {
    if (isSuccess) {
      queryClient.invalidateQueries({ queryKey: ["myJobs"] });
    }
  }, [isSuccess, queryClient]);

  return (
    <div className="mt-4 rounded-xl border border-slate-700/60 bg-slate-900/60 p-4">
      <h4 className="text-sm font-semibold text-slate-100 mb-3 flex items-center gap-2">
        <ShieldCheck className="h-4 w-4 text-slate-400" />
        Dispute Settlement
      </h4>

      <div className="mb-4 text-xs text-slate-400">
        Total Contract Balance: <span className="font-mono text-slate-200">{formatEther(totalBalance)} ETH</span>
      </div>

      <div className="space-y-3">
        <div>
          <label className="mb-1.5 flex items-center gap-1.5 text-xs font-medium text-slate-400">
            <DollarSign className="h-3 w-3" />
            Refund to Client (ETH)
          </label>
          <input
            type="number"
            step="0.01"
            min="0"
            value={clientRefund}
            onChange={(e) => setClientRefund(e.target.value)}
            className="w-full rounded-xl border border-slate-700/60 bg-slate-800/60 px-4 py-2.5 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-slate-500 focus:ring-1 focus:ring-slate-500/50"
            placeholder="0.0"
          />
        </div>
        <div>
          <label className="mb-1.5 flex items-center gap-1.5 text-xs font-medium text-slate-400">
            <DollarSign className="h-3 w-3" />
            Pay to Freelancer (ETH)
          </label>
          <input
            type="number"
            step="0.01"
            min="0"
            value={freelancerPayment}
            onChange={(e) => setFreelancerPayment(e.target.value)}
            className="w-full rounded-xl border border-slate-700/60 bg-slate-800/60 px-4 py-2.5 text-sm text-slate-200 placeholder-slate-600 outline-none transition-colors focus:border-slate-500 focus:ring-1 focus:ring-slate-500/50"
            placeholder="0.0"
          />
        </div>

        {error && (
          <p className="mt-2 text-xs text-red-400/80">
            {error.message.split("\n")[0]}
          </p>
        )}

        <button
          onClick={handleSettle}
          disabled={isPending || isConfirming || !clientRefund || !freelancerPayment}
          className="mt-2 flex w-full items-center justify-center gap-2 rounded-xl bg-slate-700 px-4 py-2.5 text-sm font-semibold text-white transition-all hover:bg-slate-600 disabled:opacity-50"
        >
          {isPending || isConfirming ? (
            <>
              <Loader2 className="h-4 w-4 animate-spin" />
              Processing...
            </>
          ) : (
            "Execute Binding Settlement"
          )}
        </button>
      </div>
    </div>
  );
}
