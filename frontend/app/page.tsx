"use client";

import Link from "next/link";
import { Lock, Zap, Scale, ArrowRight } from "lucide-react";

const features = [
  {
    icon: Lock,
    title: "Trustless Escrow",
    description:
      "Funds are locked in an on-chain smart contract. Neither party can withdraw unilaterally — trust the code, not the counterparty.",
  },
  {
    icon: Zap,
    title: "Instant Settlement",
    description:
      "Once both parties confirm completion, payment is released automatically. No middlemen, no delays, no processing fees.",
  },
  {
    icon: Scale,
    title: "Fair Arbitration",
    description:
      "Disputes are resolved on-chain with transparent evidence. Randomised arbitration ensures unbiased, tamper-proof outcomes.",
  },
];

export default function Home() {
  return (
    <div
      style={{
        minHeight: "100vh",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        background: "linear-gradient(145deg, #0a0e1a 0%, #0d1525 40%, #111b2e 100%)",
        fontFamily: "var(--font-geist-sans), system-ui, sans-serif",
        position: "relative",
        overflow: "hidden",
      }}
    >
      {/* Subtle grid overlay */}
      <div
        style={{
          position: "absolute",
          inset: 0,
          backgroundImage:
            "linear-gradient(rgba(59, 130, 246, 0.03) 1px, transparent 1px), linear-gradient(90deg, rgba(59, 130, 246, 0.03) 1px, transparent 1px)",
          backgroundSize: "64px 64px",
          pointerEvents: "none",
        }}
      />

      {/* Glow orb */}
      <div
        style={{
          position: "absolute",
          top: "15%",
          left: "50%",
          transform: "translateX(-50%)",
          width: "600px",
          height: "600px",
          background: "radial-gradient(circle, rgba(59, 130, 246, 0.08) 0%, transparent 70%)",
          borderRadius: "50%",
          pointerEvents: "none",
        }}
      />

      <main
        style={{
          position: "relative",
          zIndex: 1,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          gap: "3rem",
          padding: "2rem",
          maxWidth: "960px",
          width: "100%",
        }}
      >
        {/* Logo / Brand */}
        <div style={{ display: "flex", alignItems: "center", gap: "0.75rem" }}>
          <div
            style={{
              width: "40px",
              height: "40px",
              borderRadius: "10px",
              background: "linear-gradient(135deg, #3B82F6 0%, #1D4ED8 100%)",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              boxShadow: "0 0 24px rgba(59, 130, 246, 0.3)",
            }}
          >
            <svg
              width="22"
              height="22"
              viewBox="0 0 24 24"
              fill="none"
              stroke="white"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
            >
              <rect x="3" y="11" width="18" height="11" rx="2" ry="2" />
              <path d="M7 11V7a5 5 0 0 1 10 0v4" />
            </svg>
          </div>
          <span
            style={{
              fontSize: "1.25rem",
              fontWeight: 600,
              color: "#e2e8f0",
              letterSpacing: "-0.02em",
            }}
          >
            Freelance Escrow
          </span>
        </div>

        {/* Heading */}
        <div style={{ textAlign: "center", maxWidth: "560px" }}>
          <h1
            style={{
              fontSize: "clamp(2rem, 5vw, 3rem)",
              fontWeight: 700,
              lineHeight: 1.15,
              color: "#f1f5f9",
              letterSpacing: "-0.03em",
              margin: 0,
            }}
          >
            Decentralised
            <br />
            <span
              style={{
                background: "linear-gradient(90deg, #3B82F6, #60A5FA)",
                WebkitBackgroundClip: "text",
                WebkitTextFillColor: "transparent",
              }}
            >
              Digital Vault
            </span>
          </h1>
          <p
            style={{
              marginTop: "1rem",
              fontSize: "1.05rem",
              lineHeight: 1.7,
              color: "#94a3b8",
            }}
          >
            Trustless escrow for freelance agreements. Stake, deliver, and
            settle — enforced entirely on-chain.
          </p>
        </div>

        {/* ── 3-Column Infographic ─────────────────────────────────── */}
        <div
          style={{
            display: "grid",
            gridTemplateColumns: "repeat(3, 1fr)",
            gap: "1.25rem",
            width: "100%",
          }}
        >
          {features.map(({ icon: Icon, title, description }) => (
            <div
              key={title}
              style={{
                padding: "1.5rem",
                borderRadius: "16px",
                border: "1px solid rgba(59, 130, 246, 0.12)",
                background: "rgba(15, 23, 42, 0.5)",
                backdropFilter: "blur(12px)",
                display: "flex",
                flexDirection: "column",
                gap: "0.85rem",
                transition: "border-color 0.3s, box-shadow 0.3s",
              }}
              onMouseEnter={(e) => {
                (e.currentTarget as HTMLDivElement).style.borderColor = "rgba(59, 130, 246, 0.30)";
                (e.currentTarget as HTMLDivElement).style.boxShadow = "0 0 32px rgba(59, 130, 246, 0.08)";
              }}
              onMouseLeave={(e) => {
                (e.currentTarget as HTMLDivElement).style.borderColor = "rgba(59, 130, 246, 0.12)";
                (e.currentTarget as HTMLDivElement).style.boxShadow = "none";
              }}
            >
              <div
                style={{
                  width: "40px",
                  height: "40px",
                  borderRadius: "10px",
                  background: "rgba(59, 130, 246, 0.08)",
                  border: "1px solid rgba(59, 130, 246, 0.15)",
                  display: "flex",
                  alignItems: "center",
                  justifyContent: "center",
                }}
              >
                <Icon
                  style={{ width: "20px", height: "20px", color: "#60A5FA" }}
                  strokeWidth={1.8}
                />
              </div>
              <h3
                style={{
                  fontSize: "0.95rem",
                  fontWeight: 600,
                  color: "#e2e8f0",
                  margin: 0,
                }}
              >
                {title}
              </h3>
              <p
                style={{
                  fontSize: "0.8rem",
                  lineHeight: 1.65,
                  color: "#64748b",
                  margin: 0,
                }}
              >
                {description}
              </p>
            </div>
          ))}
        </div>

        {/* ── Enter Vault CTA ──────────────────────────────────────── */}
        <Link
          href="/login"
          style={{
            display: "inline-flex",
            alignItems: "center",
            gap: "0.5rem",
            padding: "0.75rem 2rem",
            borderRadius: "12px",
            background: "linear-gradient(135deg, #3B82F6, #2563EB)",
            color: "#ffffff",
            fontSize: "0.9rem",
            fontWeight: 600,
            letterSpacing: "-0.01em",
            textDecoration: "none",
            boxShadow: "0 0 24px rgba(59, 130, 246, 0.25), 0 4px 16px rgba(0, 0, 0, 0.3)",
            transition: "box-shadow 0.3s, transform 0.2s",
          }}
          onMouseEnter={(e) => {
            (e.currentTarget as HTMLAnchorElement).style.boxShadow =
              "0 0 36px rgba(59, 130, 246, 0.4), 0 6px 24px rgba(0, 0, 0, 0.4)";
            (e.currentTarget as HTMLAnchorElement).style.transform = "translateY(-1px)";
          }}
          onMouseLeave={(e) => {
            (e.currentTarget as HTMLAnchorElement).style.boxShadow =
              "0 0 24px rgba(59, 130, 246, 0.25), 0 4px 16px rgba(0, 0, 0, 0.3)";
            (e.currentTarget as HTMLAnchorElement).style.transform = "translateY(0)";
          }}
        >
          Enter Vault
          <ArrowRight style={{ width: "16px", height: "16px" }} />
        </Link>

        {/* Footer status */}
        <div
          style={{
            display: "flex",
            alignItems: "center",
            gap: "0.5rem",
          }}
        >
          <div
            style={{
              width: "6px",
              height: "6px",
              borderRadius: "50%",
              backgroundColor: "#22c55e",
              boxShadow: "0 0 6px #22c55e",
            }}
          />
          <span style={{ fontSize: "0.75rem", color: "#64748b" }}>
            Local Anvil Node · Chain ID 31337
          </span>
        </div>
      </main>
    </div>
  );
}
