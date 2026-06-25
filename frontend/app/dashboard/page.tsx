"use client"

import React from "react"

const Dashboard = () => {
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
      }}
    >
      <div
        style={{
          padding: "2.5rem 3rem",
          borderRadius: "16px",
          border: "1px solid rgba(59, 130, 246, 0.15)",
          background: "rgba(15, 23, 42, 0.6)",
          backdropFilter: "blur(12px)",
          textAlign: "center",
          boxShadow: "0 4px 32px rgba(0, 0, 0, 0.3)",
        }}
      >
        <h1
          style={{
            fontSize: "1.75rem",
            fontWeight: 700,
            color: "#f1f5f9",
            letterSpacing: "-0.02em",
            margin: 0,
          }}
        >
          Dashboard
        </h1>
        <p
          style={{
            marginTop: "0.75rem",
            fontSize: "0.95rem",
            color: "#94a3b8",
            lineHeight: 1.6,
          }}
        >
          Your escrow management console is being built.
        </p>
        <div
          style={{
            marginTop: "1.5rem",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            gap: "0.5rem",
          }}
        >
          <div
            style={{
              width: "6px",
              height: "6px",
              borderRadius: "50%",
              backgroundColor: "#f59e0b",
              boxShadow: "0 0 6px #f59e0b",
            }}
          />
          <span style={{ fontSize: "0.75rem", color: "#64748b" }}>
            Phase 6 · Under Construction
          </span>
        </div>
      </div>
    </div>
  )
}

export default Dashboard