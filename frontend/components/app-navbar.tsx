"use client";

import React, { useEffect, useState } from "react";
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { LayoutDashboard, Store, LogOut, Loader2 } from "lucide-react";
import { useDisconnect } from "wagmi";
import axios from "axios";

function clearFrontendCookies() {
  const cookieNames = ["username", "email", "role", "ethAccount"];
  for (const name of cookieNames) {
    document.cookie = `${name}=; Max-Age=0; path=/`;
  }
}

function getCookie(name: string): string | undefined {
  if (typeof document === "undefined") return undefined;
  const match = document.cookie.match(new RegExp(`(?:^|; )${name}=([^;]*)`));
  return match ? decodeURIComponent(match[1]) : undefined;
}

export default function AppNavbar() {
  const pathname = usePathname();
  const router = useRouter();
  const { disconnect } = useDisconnect();
  const [signingOut, setSigningOut] = useState(false);
  const [role, setRole] = useState<string>("");

  useEffect(() => {
    const init = () => setRole(getCookie("role") ?? "");
    init();
  }, []);

  const handleSignOut = async () => {
    setSigningOut(true);
    try {
      await axios.post(
        `${process.env.NEXT_PUBLIC_API_URL}/auth/logout`,
        {},
        { withCredentials: true }
      );
    } catch {
      console.warn("Backend logout request failed, clearing local state anyway.");
    }

    clearFrontendCookies();
    disconnect();
    router.push("/");
    setSigningOut(false);
  };

  const navItems = [
    { href: "/dashboard", label: "My Dashboard", icon: LayoutDashboard, show: true },
    { href: "/marketplace", label: "Marketplace", icon: Store, show: role !== "client" && role !== "arbitrator" },
  ].filter((item) => item.show);

  return (
    <nav className="sticky top-0 z-50 w-full border-b bg-[var(--vault-bg-start)]/80 backdrop-blur-xl"
      style={{ borderColor: `rgba(var(--vault-accent), 0.10)` }}
    >
      <div className="mx-auto flex h-14 max-w-6xl items-center justify-between px-6">
        {/* Brand */}
        <Link href="/" className="flex items-center gap-2.5 group">
          <div
            className="flex h-8 w-8 items-center justify-center rounded-lg transition-shadow"
            style={{
              background: `linear-gradient(135deg, rgba(var(--vault-accent), 0.9), rgba(var(--vault-accent), 0.6))`,
              boxShadow: `0 0 16px rgba(var(--vault-accent), 0.25)`,
            }}
          >
            <svg
              width="16"
              height="16"
              viewBox="0 0 24 24"
              fill="none"
              stroke="white"
              strokeWidth="2.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            >
              <rect x="3" y="11" width="18" height="11" rx="2" ry="2" />
              <path d="M7 11V7a5 5 0 0 1 10 0v4" />
            </svg>
          </div>
          <span className="text-sm font-semibold tracking-tight text-slate-200">
            Freelance Escrow
          </span>
        </Link>

        {/* Nav Links */}
        <div className="flex items-center gap-1">
          {navItems.map(({ href, label, icon: Icon }) => {
            const isActive = pathname === href;
            return (
              <Link
                key={href}
                href={href}
                className={`flex items-center gap-2 rounded-lg px-3.5 py-2 text-sm font-medium transition-all duration-200 ${
                  isActive
                    ? "text-slate-200"
                    : "text-slate-400 hover:bg-white/[0.04] hover:text-slate-200"
                }`}
                style={
                  isActive
                    ? {
                        background: `rgba(var(--vault-accent), 0.10)`,
                        boxShadow: `inset 0 0 0 1px rgba(var(--vault-accent), 0.20)`,
                        color: `rgba(var(--vault-accent), 1)`,
                      }
                    : undefined
                }
              >
                <Icon className="h-4 w-4" />
                {label}
              </Link>
            );
          })}
        </div>

        {/* Actions */}
        <div className="flex items-center gap-2">
          {role && (
            <span
              className="inline-flex items-center gap-1.5 rounded-full border border-slate-700/50 bg-slate-800/60 px-2.5 py-1 text-[10px] font-semibold uppercase tracking-wider text-slate-400"
            >
              <span
                className="h-1.5 w-1.5 rounded-full"
                style={{ background: `rgba(var(--vault-accent), 1)` }}
              />
              {role}
            </span>
          )}

          <button
            onClick={handleSignOut}
            disabled={signingOut}
            className="flex items-center gap-1.5 rounded-lg border border-red-500/20 bg-red-500/[0.06] px-3 py-1.5 text-xs font-medium text-red-400 transition-all duration-200 hover:border-red-500/40 hover:bg-red-500/[0.12] hover:text-red-300 hover:shadow-[0_0_12px_rgba(239,68,68,0.15)] disabled:pointer-events-none disabled:opacity-50"
          >
            {signingOut ? (
              <Loader2 className="h-3.5 w-3.5 animate-spin" />
            ) : (
              <LogOut className="h-3.5 w-3.5" />
            )}
            Sign Out
          </button>
        </div>
      </div>
    </nav>
  );
}
