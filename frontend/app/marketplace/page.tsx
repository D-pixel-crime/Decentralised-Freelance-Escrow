"use client";

import React, { useEffect } from "react";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import { useRouter } from "next/navigation";
import Link from "next/link";

import type { Job } from "@/types/job";
import {
  Globe, ExternalLink, Loader2, PackageOpen,
  DollarSign, Calendar,
} from "lucide-react";
import AppNavbar from "@/components/app-navbar";

function truncateAddress(addr: string) {
  if (!addr || addr.length < 12) return addr || "—";
  return `${addr.slice(0, 6)}…${addr.slice(-4)}`;
}

function getCookie(name: string): string | undefined {
  if (typeof document === "undefined") return undefined;
  const match = document.cookie.match(new RegExp(`(?:^|; )${name}=([^;]*)`));
  return match ? decodeURIComponent(match[1]) : undefined;
}

// ── Marketplace page ────────────────────────────────────────────────────────

export default function MarketplacePage() {
  const router = useRouter();

  useEffect(() => {
    const role = getCookie("role");
    if (role === "client" || role === "arbitrator") {
      router.replace("/dashboard");
    }
  }, [router]);

  const { data: jobs, isLoading, isError } = useQuery<Job[]>({
    queryKey: ["openJobs"],
    queryFn: async () => {
      const res = await axios.get(
        `${process.env.NEXT_PUBLIC_API_URL}/api/get/jobs/open`,
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
              <Globe className="h-5 w-5" style={{ color: `rgba(var(--vault-accent), 1)` }} />
            </div>
            <h1 className="text-2xl font-bold tracking-tight text-slate-100">
              Marketplace
            </h1>
          </div>
          <p className="text-sm text-slate-500 pl-[52px]">
            Browse open escrow jobs looking for freelancers
          </p>
        </div>

        {/* Loading */}
        {isLoading && (
          <div className="flex flex-col items-center justify-center gap-3 py-32">
            <Loader2 className="h-8 w-8 animate-spin" style={{ color: `rgba(var(--vault-accent), 0.6)` }} />
            <span className="text-sm text-slate-500">Loading marketplace…</span>
          </div>
        )}

        {/* Error */}
        {isError && (
          <div className="flex flex-col items-center justify-center gap-3 py-32">
            <div className="flex h-14 w-14 items-center justify-center rounded-full bg-red-500/10 border border-red-500/20">
              <span className="text-2xl">⚠</span>
            </div>
            <p className="text-sm text-slate-400">
              Failed to load open jobs. Please try again later.
            </p>
          </div>
        )}

        {/* Empty */}
        {!isLoading && !isError && uniqueJobs.length === 0 && (
          <div className="flex flex-col items-center justify-center gap-4 py-32">
            <div className="flex h-16 w-16 items-center justify-center rounded-2xl bg-slate-800/60 border border-slate-700/40">
              <PackageOpen className="h-7 w-7 text-slate-600" />
            </div>
            <div className="text-center">
              <p className="text-sm font-medium text-slate-400">
                No open jobs available
              </p>
              <p className="mt-1 text-xs text-slate-600">
                Check back later — new escrow agreements appear here once posted.
              </p>
            </div>
          </div>
        )}

        {/* Job Cards */}
        {!isLoading && !isError && uniqueJobs.length > 0 && (
          <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
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
                {/* Status + Title */}
                <div className="mb-3">
                  <span
                    className="inline-flex items-center gap-1.5 rounded-full px-2.5 py-1 text-xs font-medium mb-2"
                    style={{
                      background: `rgba(var(--vault-accent), 0.10)`,
                      color: `rgba(var(--vault-accent), 1)`,
                    }}
                  >
                    <span
                      className="h-1.5 w-1.5 rounded-full animate-pulse"
                      style={{ background: `rgba(var(--vault-accent), 1)` }}
                    />
                    OPEN
                  </span>
                  <h3 className="mt-2 text-base font-semibold text-slate-100 leading-snug">
                    {job.title || "Untitled Job"}
                  </h3>
                </div>

                {/* Description */}
                {job.description && (
                  <p className="mb-3 text-xs text-slate-400 leading-relaxed line-clamp-3">
                    {job.description}
                  </p>
                )}

                {/* Pay Range & Deadline pills */}
                <div className="mb-4 flex flex-wrap gap-2">
                  {(job.payMin !== undefined || job.payMax !== undefined) && (
                    <span className="inline-flex items-center gap-1.5 rounded-lg bg-emerald-500/10 px-2.5 py-1 text-xs font-medium text-emerald-400">
                      <DollarSign className="h-3.5 w-3.5" />
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

                {/* Job ID */}
                <div className="mb-4">
                  <span className="text-[10px] font-medium uppercase tracking-widest text-slate-600">
                    Job ID
                  </span>
                  <p className="mt-1 font-mono text-sm text-slate-300">
                    {truncateAddress(job.id)}
                  </p>
                </div>

                {/* Client */}
                <div className="border-t border-slate-800/60 pt-4">
                  <span className="text-[10px] font-medium uppercase tracking-widest text-slate-600">
                    Posted by
                  </span>
                  <div className="mt-1 flex items-center gap-2">
                    <p className="font-mono text-xs text-slate-400">
                      {truncateAddress(job.clientId)}
                    </p>
                    <ExternalLink className="h-3 w-3 text-slate-600 opacity-0 transition-opacity group-hover:opacity-100" />
                  </div>
                </div>

                <div className="mt-4">
                  <Link
                    href={`/marketplace/${job.id}`}
                    className="flex w-full items-center justify-center gap-2 rounded-xl px-4 py-2.5 text-sm font-semibold text-white transition-all duration-200 hover:bg-slate-800"
                    style={{
                      background: `rgba(var(--vault-accent), 0.85)`,
                      boxShadow: `0 0 20px rgba(var(--vault-accent), 0.20)`,
                    }}
                  >
                    View Details
                  </Link>
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
      </main>
    </div>
  );
}
