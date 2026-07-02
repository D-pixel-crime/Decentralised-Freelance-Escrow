"use client";

import { useEffect, useRef } from "react";
import { useWriteContract, useWaitForTransactionReceipt } from "wagmi";
import { useQueryClient } from "@tanstack/react-query";
import { FREELANCE_ESCROW_ABI } from "@/constants/contract";

/**
 * useEscrowAction — DRY hook for non-payable escrow lifecycle calls.
 *
 * Wraps `useWriteContract` → `useWaitForTransactionReceipt` → auto-invalidate.
 * Reports busy state changes via optional `onBusyChange` callback.
 *
 * SELF-CORRECTION AUDIT: All lifecycle functions (requestPayment, acceptJobCompletion,
 * rejectJobCompletion, breakDeal, cancelDealBreak, raiseDispute) are `nonpayable`.
 * This hook NEVER sends a `value` field.
 */

/** Non-payable lifecycle function names from the Escrow ABI */
type EscrowLifecycleFn =
  | "requestPayment"
  | "acceptJobCompletion"
  | "rejectJobCompletion"
  | "breakDeal"
  | "cancelDealBreak"
  | "raiseDispute";

export function useEscrowAction(
  contractAddress: string,
  functionName: EscrowLifecycleFn,
  onBusyChange?: (busy: boolean) => void
) {
  const queryClient = useQueryClient();

  const {
    data: hash,
    writeContract,
    isPending,
    error: writeError,
    reset,
  } = useWriteContract();

  const {
    isLoading: isConfirming,
    isSuccess,
    error: receiptError,
  } = useWaitForTransactionReceipt({ hash });

  // ── Auto-invalidate on success ──────────────────────────────────────
  const hasInvalidated = useRef(false);

  useEffect(() => {
    if (isSuccess && !hasInvalidated.current) {
      hasInvalidated.current = true;
      queryClient.invalidateQueries({ queryKey: ["myJobs"] });
    }
  }, [isSuccess, queryClient]);

  // Reset the invalidation guard when hash changes (new tx)
  useEffect(() => {
    hasInvalidated.current = false;
  }, [hash]);

  // ── Report busy state upstream ──────────────────────────────────────
  const isBusy = isPending || isConfirming;

  useEffect(() => {
    onBusyChange?.(isBusy);
  }, [isBusy, onBusyChange]);

  // ── Execute ─────────────────────────────────────────────────────────
  const execute = () => {
    reset(); // clear previous state
    writeContract({
      abi: FREELANCE_ESCROW_ABI,
      address: contractAddress as `0x${string}`,
      functionName,
      // NO value — nonpayable
    });
  };

  return {
    execute,
    isPending,
    isConfirming,
    isSuccess,
    isBusy,
    error: writeError || receiptError,
  };
}
