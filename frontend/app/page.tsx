"use client";

import { ConnectButton } from "@rainbow-me/rainbowkit";

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
          top: "20%",
          left: "50%",
          transform: "translateX(-50%)",
          width: "500px",
          height: "500px",
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
          gap: "2.5rem",
          padding: "2rem",
        }}
      >
        {/* Logo / Brand */}
        <div
          style={{
            display: "flex",
            alignItems: "center",
            gap: "0.75rem",
          }}
        >
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

        {/* Connect Button */}
        <div
          style={{
            padding: "1.5rem 2rem",
            borderRadius: "16px",
            border: "1px solid rgba(59, 130, 246, 0.15)",
            background: "rgba(15, 23, 42, 0.6)",
            backdropFilter: "blur(12px)",
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            gap: "1rem",
            boxShadow: "0 4px 32px rgba(0, 0, 0, 0.3)",
          }}
        >
          <span
            style={{
              fontSize: "0.8rem",
              fontWeight: 500,
              color: "#64748b",
              textTransform: "uppercase",
              letterSpacing: "0.1em",
            }}
          >
            Connect to Anvil
          </span>
          <ConnectButton />
        </div>

        {/* Footer status */}
        <div
          style={{
            display: "flex",
            alignItems: "center",
            gap: "0.5rem",
            marginTop: "1rem",
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
