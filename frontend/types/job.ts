export type JobStatus =
  | "UNALLOCATED"
  | "AGREED"
  | "CLIENT_STAKED"
  | "FREELANCER_STAKED"
  | "ALL_STAKED_AND_PENDING"
  | "PENDING_CLIENT_CONFIRMATION"
  | "JOB_COMPLETED"
  | "CANCEL_REQUESTED"
  | "DEAL_BROKEN"
  | "RANDOM_DISPUTED"
  | "PAYMENT_DISPUTED";

export interface Job {
  id: string;
  contractAddress: string;
  clientId: string;
  freelancerId: string;
  status: JobStatus;
  applicants?: string[];
  deliverableCid?: string;
  arbitratorEth?: string;

  // Web2 metadata (off-chain)
  title?: string;
  description?: string;
  deadline?: string;
  contactEmail?: string;
  payMin?: number;
  payMax?: number;
}
