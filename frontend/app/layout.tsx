import type { Metadata } from "next";
import { Courier_Prime, Geist } from "next/font/google";
import "./globals.css";

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
      className={`${geist.className} ${courierPrime.variable} h-full antialiased`}
    >
      <body className="min-h-screen w-full flex-col]">
        {children}
      </body>
    </html>
  );
}
