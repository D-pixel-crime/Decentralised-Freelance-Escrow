"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { WagmiProvider, http } from "wagmi";
import { foundry } from "wagmi/chains";
import { getDefaultConfig, RainbowKitProvider, darkTheme } from "@rainbow-me/rainbowkit";
import "@rainbow-me/rainbowkit/styles.css";

const config = getDefaultConfig({
  appName: "Freelance Escrow",
  projectId: "00000000000000000000000000000000", // placeholder — required by WalletConnect but unused for local Anvil
  chains: [foundry],
  transports: {
    [foundry.id]: http("http://127.0.0.1:8545"),
  },
  ssr: true,
});

const queryClient = new QueryClient();

export default function Web3Provider({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        <RainbowKitProvider
          theme={darkTheme({
            accentColor: "#3B82F6",
            accentColorForeground: "#FFFFFF",
            borderRadius: "medium",
            fontStack: "system",
          })}
        >
          {children}
        </RainbowKitProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}
