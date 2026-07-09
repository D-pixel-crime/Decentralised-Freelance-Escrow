"use client";

import { useEffect, useRef } from "react";
import { useWriteContract, useWaitForTransactionReceipt } from "wagmi";
import { useQueryClient } from "@tanstack/react-query";
import { FREELANCE_ESCROW_ABI } from "@/constants/contract";
import { useToast } from "@/contexts/ToastContext";

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
  const toast = useToast();

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

  const execute = () => {
    reset(); // clear previous state
    writeContract({
      abi: FREELANCE_ESCROW_ABI,
      address: contractAddress as `0x${string}`,
      functionName,
      // NO value — nonpayable
    }, {
      onError: (err) => {
        let msg = (err as Error & { shortMessage?: string }).shortMessage || err.message || "Transaction failed";
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

  return {
    execute,
    isPending,
    isConfirming,
    isSuccess,
    isBusy,
    error: writeError || receiptError,
  };
}
