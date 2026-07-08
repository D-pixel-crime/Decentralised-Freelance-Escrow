'use client'

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import Link from "next/link"
import React, { useState } from "react"
import axios from "axios";
import { getAddress } from "ethers";
import { useRouter } from "next/navigation";
import { ConnectButton } from "@rainbow-me/rainbowkit";
import { useAccount } from "wagmi";
import { useToast } from "@/contexts/ToastContext";
import { extractErrorMsg } from "@/lib/utils";

const SignupForm = ({ ...props }: React.ComponentProps<typeof Card>) => {
  const [details, setDetails] = useState({ email: "", username: "", role: "" })
  const toast = useToast();
  const router = useRouter();
  const { address, isConnected } = useAccount();

  // Derive the checksummed wallet address from Wagmi's connected account
  const walletAddr = isConnected && address ? getAddress(address) : "";

  const handleSignup = async (e: React.SubmitEvent) => {
    e.preventDefault()
    if (!walletAddr) {
      toast.error("Please connect your wallet using the button above.");
      return;
    }
    const backendUri = process.env.NEXT_PUBLIC_API_URL;
    if (!backendUri) {
      toast.error("Backend URI Missing!");
      return;
    }

    try {
      await axios.post(`${backendUri}/auth/signup`, { ...details, ethAccount: walletAddr }, { withCredentials: true });

      const cookieSettings = `; path=/; max-age=${3600 * 24}; SameSite=Lax`; // 1 day
      document.cookie = `username=${details.username}${cookieSettings}`;
      document.cookie = `email=${details.email}${cookieSettings}`;
      document.cookie = `role=${details.role}${cookieSettings}`;
      document.cookie = `ethAccount=${walletAddr}${cookieSettings}`;

      router.push("/dashboard");
      router.refresh();
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } catch (error: any) {
      const errorMsg = extractErrorMsg(error, "Unknown error occurred");
      toast.error("Signup failed: " + errorMsg);
      return;
    }

    setDetails({ username: "", role: "", email: "" })
  }

  return (
    <Card {...props} className="border-blue-500/15 bg-[#0f172a]/80 backdrop-blur-xl shadow-2xl shadow-blue-500/5">
      <CardHeader className="flex-center w-full flex-col">
        <CardTitle className="text-lg text-slate-100">Create an account</CardTitle>
        <CardDescription className="text-center text-slate-400">
          Enter your information and connect your wallet to get started
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSignup}>
          <FieldGroup>
            <Field>
              <FieldLabel htmlFor="username">Username<span className="text-destructive">*</span></FieldLabel>
              <Input id="username" type="text" placeholder="john_doe_69" required value={details.username} onChange={(e) => setDetails({ ...details, username: e.target.value })} />
            </Field>
            <Field>
              <FieldLabel htmlFor="email">Email<span className="text-destructive">*</span></FieldLabel>
              <Input
                id="email"
                type="email"
                placeholder="m@example.com"
                required
                value={details.email} onChange={(e) => setDetails({ ...details, email: e.target.value })}
              />
              <FieldDescription>
                We&apos;ll use this to contact you. We will not share your email with anyone else.
              </FieldDescription>
            </Field>
            <Field>
              <FieldLabel htmlFor="role">Select Role<span className="text-destructive">*</span></FieldLabel>
              <Select value={details.role} onValueChange={(val) => setDetails({ ...details, role: val })} required>
                <SelectTrigger>
                  <SelectValue placeholder="Select a role" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    <SelectLabel>Role</SelectLabel>
                    <SelectItem value="freelancer">Freelancer</SelectItem>
                    <SelectItem value="client">Client</SelectItem>
                    <SelectItem value="arbitrator">Arbitrator</SelectItem>
                  </SelectGroup>
                </SelectContent>
              </Select>
            </Field>
            <Field>
              <FieldLabel htmlFor="ethAccount">Ethereum Wallet<span className="text-destructive">*</span></FieldLabel>
              <div className="flex flex-col gap-3">
                <div className="flex items-center justify-center rounded-xl border border-slate-700/50 bg-slate-800/40 p-3">
                  <ConnectButton />
                </div>
                {walletAddr && (
                  <Input
                    id="ethAccount"
                    value={walletAddr}
                    type="text"
                    disabled
                    className="disabled:bg-slate-800/50 disabled:text-slate-400 disabled:cursor-not-allowed disabled:border-slate-700/50 font-mono text-xs"
                  />
                )}
              </div>
            </Field>
            <FieldGroup>
              <Field>
                <Button type="submit" disabled={!isConnected} className="cursor-pointer bg-blue-600 hover:bg-blue-500 text-white w-full disabled:opacity-50 disabled:cursor-not-allowed">
                  {isConnected ? "Create Account" : "Connect Wallet First"}
                </Button>
                <FieldDescription className="px-6 text-center text-slate-500">
                  Already have an account? <Link href="/login" className="text-blue-400 hover:text-blue-300 transition-colors">Login</Link>
                </FieldDescription>
              </Field>
            </FieldGroup>
          </FieldGroup>
        </form>
      </CardContent>
    </Card>
  )
}

export default SignupForm