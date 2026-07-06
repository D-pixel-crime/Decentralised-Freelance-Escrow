"use client";

import React, { useEffect, useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { useRouter } from "next/navigation";
import { useAccount } from "wagmi";
import type { Job } from "@/types/job";
import {
  Globe, ExternalLink, Loader2, DollarSign, Calendar, FileText, CheckCircle
} from "lucide-react";
import AppNavbar from "@/components/app-navbar";

function truncateAddress(addr: string) {
  if (!addr || addr.length < 12) return addr || "—";
  return `${addr.slice(0, 6)}…${addr.slice(-4)}`;
}

export default function JobDetailsPage({ params }: { params: any }) {
  const router = useRouter();
  const { address } = useAccount();
  const queryClient = useQueryClient();
  
  const unwrappedParams = React.use(params as Promise<{ id: string }>);
  const jobId = unwrappedParams.id;

  const { data: job, isLoading, isError } = useQuery<Job>({
    queryKey: ["job", jobId],
    queryFn: async () => {
      const res = await axios.get(
        `${process.env.NEXT_PUBLIC_API_URL}/api/get/job/${jobId}`,
        { withCredentials: true }
      );
      return res.data;
    },
    enabled: !!jobId,
  });

  const mutation = useMutation({
    mutationFn: async () => {
      await axios.post(
        `${process.env.NEXT_PUBLIC_API_URL}/api/post/job/apply`,
        { jobId },
        { withCredentials: true }
      );
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["job", jobId] });
    },
  });

  if (isLoading) {
    return (
      <div className="min-h-screen bg-slate-950">
        <AppNavbar />
        <div className="flex flex-col items-center justify-center gap-3 py-32">
          <Loader2 className="h-8 w-8 animate-spin" style={{ color: `rgba(var(--vault-accent), 0.6)` }} />
          <span className="text-sm text-slate-500">Loading job details…</span>
        </div>
      </div>
    );
  }

  if (isError || !job) {
    return (
      <div className="min-h-screen bg-slate-950">
        <AppNavbar />
        <div className="flex flex-col items-center justify-center gap-3 py-32">
          <div className="flex h-14 w-14 items-center justify-center rounded-full bg-red-500/10 border border-red-500/20">
            <span className="text-2xl">⚠</span>
          </div>
          <p className="text-sm text-slate-400">Failed to load job details. Please try again later.</p>
          <button onClick={() => router.back()} className="mt-4 text-sm text-slate-500 hover:text-slate-300">
            Go Back
          </button>
        </div>
      </div>
    );
  }

  const hasApplied = address && job.applicants?.includes(address);

  return (
    <div
      className="min-h-screen font-[family-name:var(--font-geist-sans)]"
      style={{
        background: `linear-gradient(to bottom, var(--vault-bg-start), var(--vault-bg-mid), var(--vault-bg-end))`,
      }}
    >
      <AppNavbar />

      <div
        className="pointer-events-none fixed inset-0"
        style={{
          backgroundImage:
            `linear-gradient(var(--vault-grid) 1px, transparent 1px), linear-gradient(90deg, var(--vault-grid) 1px, transparent 1px)`,
          backgroundSize: "64px 64px",
        }}
      />

      <main className="relative z-10 mx-auto max-w-4xl px-6 py-10">
        <button onClick={() => router.push("/marketplace")} className="mb-6 text-sm text-slate-400 hover:text-slate-200 flex items-center gap-2">
          ← Back to Marketplace
        </button>

        <div
          className="rounded-2xl backdrop-blur-sm p-8"
          style={{
            border: `1px solid rgba(var(--vault-accent), 0.15)`,
            background: `linear-gradient(145deg, rgba(15, 23, 42, 0.8), rgba(10, 14, 26, 0.9))`,
          }}
        >
          <div className="mb-6 flex flex-wrap gap-3">
            <span
              className="inline-flex items-center gap-1.5 rounded-full px-3 py-1 text-xs font-medium"
              style={{
                background: `rgba(var(--vault-accent), 0.10)`,
                color: `rgba(var(--vault-accent), 1)`,
              }}
            >
              <span
                className="h-1.5 w-1.5 rounded-full animate-pulse"
                style={{ background: `rgba(var(--vault-accent), 1)` }}
              />
              {job.status}
            </span>
            {(job.payMin !== undefined || job.payMax !== undefined) && (
              <span className="inline-flex items-center gap-1 rounded-full bg-emerald-500/10 px-3 py-1 text-xs font-medium text-emerald-400 border border-emerald-500/20">
                <DollarSign className="h-3 w-3" />
                {job.payMin === job.payMax ? `${job.payMin} ETH` : `${job.payMin} - ${job.payMax} ETH`}
              </span>
            )}
            {job.deadline && (
              <span className="inline-flex items-center gap-1 rounded-full bg-blue-500/10 px-3 py-1 text-xs font-medium text-blue-400 border border-blue-500/20">
                <Calendar className="h-3 w-3" />
                {job.deadline}
              </span>
            )}
          </div>

          <h1 className="text-3xl font-bold text-slate-100 mb-6">{job.title || "Untitled Job"}</h1>

          <div className="grid md:grid-cols-2 gap-8 mb-8 border-t border-b border-slate-800/60 py-6">
            <div>
              <h3 className="text-sm font-semibold text-slate-300 mb-2">Job Description</h3>
              <p className="text-sm text-slate-400 whitespace-pre-wrap leading-relaxed">
                {job.description || "No description provided."}
              </p>
            </div>
            
            <div className="space-y-4">
              <div>
                <span className="text-[10px] font-medium uppercase tracking-widest text-slate-600">Client Address</span>
                <p className="mt-1 font-mono text-sm text-slate-300">{job.clientId}</p>
              </div>
              <div>
                <span className="text-[10px] font-medium uppercase tracking-widest text-slate-600">Contact Email</span>
                <p className="mt-1 text-sm text-slate-300 italic">Masked (Revealed during dispute)</p>
              </div>
              <div>
                <span className="text-[10px] font-medium uppercase tracking-widest text-slate-600">Applicants</span>
                <p className="mt-1 text-sm text-slate-300">{job.applicants?.length || 0} freelancer(s) applied</p>
              </div>
            </div>
          </div>

          <div className="mt-8 flex justify-end">
            {!address ? (
              <p className="text-sm text-slate-500">Please connect your wallet to apply.</p>
            ) : hasApplied ? (
              <div className="flex items-center gap-2 rounded-xl border border-emerald-500/20 bg-emerald-500/10 px-6 py-3">
                <CheckCircle className="h-5 w-5 text-emerald-400" />
                <span className="text-sm font-semibold text-emerald-400">Already Applied</span>
              </div>
            ) : (
              <button
                onClick={() => mutation.mutate()}
                disabled={mutation.isPending}
                className="flex items-center gap-2 rounded-xl px-8 py-3 text-sm font-semibold text-white transition-all duration-200 disabled:cursor-not-allowed disabled:opacity-50"
                style={{
                  background: `rgba(var(--vault-accent), 0.85)`,
                  boxShadow: `0 0 24px rgba(var(--vault-accent), 0.20)`,
                }}
              >
                {mutation.isPending ? (
                  <>
                    <Loader2 className="h-5 w-5 animate-spin" />
                    Applying...
                  </>
                ) : (
                  <>
                    <FileText className="h-5 w-5" />
                    Apply for Job
                  </>
                )}
              </button>
            )}
          </div>
          
          {mutation.isError && (
             <div className="mt-4 rounded-xl border border-red-500/20 bg-red-500/10 px-4 py-3">
               <p className="text-sm text-red-400">
                 Failed to apply. You might not be registered as a freelancer.
               </p>
             </div>
          )}
        </div>
      </main>
    </div>
  );
}
