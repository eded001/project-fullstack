import type { Metadata } from "next";
import { Courier_Prime, Geist, Inter } from "next/font/google";
import "./globals.css";
import { cn } from "@/lib/utils";

const inter = Inter({subsets:['latin'],variable:'--font-sans'});

const geist = Geist({ subsets: ["latin"] });

const courierPrime = Courier_Prime({
  subsets: ["latin"],
  weight: ["400", "700"],
  style: ["normal", "italic"],
  variable: "--font-courier-prime",
});

export const metadata: Metadata = {
  title: "system-jav",
  description: "Projeto fullstack com Next.js e Go.",
};

export default function RootLayout({ children }: LayoutProps<"/">) {
  return (
    <html
      lang="pt-BR"
      className={cn("h-full", "antialiased", geist.className, courierPrime.variable, "font-sans", inter.variable)}
    >
      <body className="min-h-screen w-full flex-col]">
        {children}
      </body>
    </html>
  );
}
