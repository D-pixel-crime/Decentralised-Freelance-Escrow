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

const SignupForm = ({ ...props }: React.ComponentProps<typeof Card>) => {
  const [walletAddr, setWalletAddr] = useState("");
  const [details, setDetails] = useState({ email: "", username: "", role: "" })

  const router = useRouter();

  const connectToWallet = async () => {
    if (typeof window.ethereum !== 'undefined') {
      try {
        const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
        const ethAccount = getAddress(accounts[0]);
        setWalletAddr(ethAccount);
      } catch (error) {
        alert("Error fetching wallet details! Please try again later.");
        console.log(error);
      }
    } else {
      alert("Please Install Metamask!");
    }
  }

  const handleSignup = async (e: React.SubmitEvent) => {
    e.preventDefault()
    if (!walletAddr) {
      alert("Please Connect your Ethereum Wallet!");
      return;
    }
    const backendUri = process.env.NEXT_PUBLIC_BACKEND_URI;
    if (!backendUri) {
      alert("Backend URI Missing!");
      return;
    }
    console.log(details);
    console.log(walletAddr);

    try {
      await axios.post(`${backendUri}/auth/signup`, { ...details, ethAccount: walletAddr });

      const cookieSettings = `; path=/; max-age=${3600 * 24}; SameSite=Lax`; // 7 days
      document.cookie = `username=${details.username}${cookieSettings}`;
      document.cookie = `email=${details.email}${cookieSettings}`;
      document.cookie = `role=${details.role}${cookieSettings}`;
      document.cookie = `ethAccount=${walletAddr}${cookieSettings}`;

      router.push("/");
      router.refresh();
    } catch (error) {
      alert("Error in Signup!" + error);
      return;
    }

    setWalletAddr("");
    setDetails({ username: "", role: "", email: "" })
  }

  return (
    <Card {...props} className="border-blue-500/15 bg-[#0f172a]/80 backdrop-blur-xl shadow-2xl shadow-blue-500/5">
      <CardHeader className="flex-center w-full flex-col">
        <CardTitle className="text-lg text-slate-100">Create an account</CardTitle>
        <CardDescription className="text-center text-slate-400">
          Enter your information below to create your account
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
                  </SelectGroup>
                </SelectContent>
              </Select>
            </Field>
            <Field>
              <FieldLabel htmlFor="ethAccount">Ethereum Wallet Address<span className="text-destructive">*</span></FieldLabel>
              <div className="flex flex-row justify-between items-center gap-2">
                <Input id="ethAccount" value={walletAddr} type="text" placeholder="0xxxxxxxxxxxxxxxxxx......" disabled className="disabled:bg-slate-800/50 disabled:text-slate-500 disabled:cursor-not-allowed disabled:border-slate-700/50" />
                <Button type="button" className="cursor-pointer bg-blue-600 hover:bg-blue-500 text-white" onClick={connectToWallet}>Connect Wallet</Button>
              </div>
            </Field>
            <FieldGroup>
              <Field>
                <Button type="submit" className="cursor-pointer bg-blue-600 hover:bg-blue-500 text-white w-full">Create Account</Button>
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