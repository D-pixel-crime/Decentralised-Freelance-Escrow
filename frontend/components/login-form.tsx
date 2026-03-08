'use client'

declare global {
    interface Window {
        ethereum?: {
            request: (args: { method: string }) => Promise<string[]>;
        };
    }
}

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

            const nonce = (await axios.post(`${process.env.NEXT_PUBLIC_BACKEND_URI}/auth/login/initiate`, { role, ethAccount: walletAddr })).data.nonce
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

            const response = await axios.post(`${process.env.NEXT_PUBLIC_BACKEND_URI}/auth/login/verify`, { ethAccount: walletAddr, role, message, signature }, { withCredentials: true })
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
        <Card {...props}>
            <CardHeader className="flex-center w-full flex-col">
                <CardTitle className="text-lg">Login to an existing account</CardTitle>
                <CardDescription className="text-center">
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
                                <Input id="ethAccount" value={walletAddr} type="text" placeholder="0xxxxxxxxxxxxxxxxxx......" disabled className="disabled:bg-gray-200 disabled:text-gray-500 disabled:cursor-not-allowed" />
                                <Button type="button" className="cursor-pointer" onClick={connectToWallet}>Connect Wallet</Button>
                            </div>
                        </Field>
                        <FieldGroup>
                            <Field>
                                <Button type="submit" className="cursor-pointer bg-green-600 hover:bg-green-500">Login</Button>
                                <FieldDescription className="px-6 text-center">
                                    Don&apos;t have an account? <Link href="/signup">Sign-up</Link>
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