"use client";

import { Geist, Geist_Mono } from "next/font/google";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import Link from "next/link";

import "./globals.scss";
import "./page.scss";

export const queryClient = new QueryClient();

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

// export const metadata: Metadata = {
//   title: "Create Next App",
//   description: "Generated by create next app",
// };

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${geistSans.variable} ${geistMono.variable}`}>
        <nav>
          <Link href={"/"}>Home</Link>
          <Link href={"/addNewWord"}>Add new</Link>
        </nav>
        <QueryClientProvider client={queryClient}>
          {children}
        </QueryClientProvider>
      </body>
    </html>
  );
}