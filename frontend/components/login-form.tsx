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
import { SiweMessage } from "siwe";
import { BrowserProvider, getAddress } from "ethers"
import { useRouter } from "next/navigation";

const LoginForm = ({ ...props }: React.ComponentProps<typeof Card>) => {
    const [walletAddr, setWalletAddr] = useState("");
    const [role, setRole] = useState("")
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

    const handleLogin = async (e: React.SubmitEvent) => {
        e.preventDefault()

        try {
            if (!window.ethereum) {
                alert("Please Connect your Ethereum Wallet!");
                return;
            }

            const nonce = (await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/auth/login/initiate`, { role, ethAccount: walletAddr }, { withCredentials: true })).data.nonce
            const provider = new BrowserProvider(window.ethereum)

            const message = new SiweMessage({
                domain: window.location.host,
                address: walletAddr,
                statement: "Sign in to Freelance Escrow.",
                uri: window.location.origin,
                version: '1',
                chainId: Number((await provider.getNetwork()).chainId),
                nonce: nonce
            }).prepareMessage();

            const signer = await provider.getSigner()
            const signature = await signer.signMessage(message);

            const response = await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/auth/login/verify`, { ethAccount: walletAddr, role, message, signature }, { withCredentials: true })
            const { username, email } = response.data;

            const cookieSettings = `; path=/; max-age=${3600 * 24}; SameSite=Lax`; // 7 days
            document.cookie = `username=${username}${cookieSettings}`;
            document.cookie = `email=${email}${cookieSettings}`;
            document.cookie = `role=${role}${cookieSettings}`;
            document.cookie = `ethAccount=${walletAddr}${cookieSettings}`;

            router.push("/");
            router.refresh();
        } catch (error) {
            alert("Error while Login.\n" + error)
            return
        }

        setWalletAddr("");
        setRole("")
    }

    return (
        <Card {...props} className="border-blue-500/15 bg-[#0f172a]/80 backdrop-blur-xl shadow-2xl shadow-blue-500/5">
            <CardHeader className="flex-center w-full flex-col">
                <CardTitle className="text-lg text-slate-100">Login to an existing account</CardTitle>
                <CardDescription className="text-center text-slate-400">
                    Enter your credentials below to login into your account
                </CardDescription>
            </CardHeader>
            <CardContent>
                <form onSubmit={handleLogin}>
                    <FieldGroup>
                        <Field>
                            <FieldLabel htmlFor="role">Select Role<span className="text-destructive">*</span></FieldLabel>
                            <Select value={role} onValueChange={(val) => setRole(val)} required>
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
                                <Button type="submit" className="cursor-pointer bg-blue-600 hover:bg-blue-500 text-white w-full">Login</Button>
                                <FieldDescription className="px-6 text-center text-slate-500">
                                    Don&apos;t have an account? <Link href="/signup" className="text-blue-400 hover:text-blue-300 transition-colors">Sign-up</Link>
                                </FieldDescription>
                            </Field>
                        </FieldGroup>
                    </FieldGroup>
                </form>
            </CardContent>
        </Card>
    )
}

export default LoginForm