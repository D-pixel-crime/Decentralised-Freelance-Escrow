import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import { cookies } from "next/headers";
import "./globals.css";
import Web3Provider from "./providers/Web3Provider";
import { ToastProvider } from "@/contexts/ToastContext";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Freelance Escrow | Decentralised Digital Vault",
  description:
    "Institutional-grade decentralised escrow for freelance agreements, powered by smart contracts.",
};

function getThemeClass(role: string | undefined): string {
  switch (role) {
    case "client":
      return "theme-client";
    case "freelancer":
      return "theme-freelancer";
    default:
      return "";
  }
}

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const cookieStore = await cookies();
  const role = cookieStore.get("role")?.value;
  const themeClass = getThemeClass(role);

  return (
    <html lang="en" className="dark">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased ${themeClass}`}
      >
        <ToastProvider>
          <Web3Provider>{children}</Web3Provider>
        </ToastProvider>
      </body>
    </html>
  );
}
