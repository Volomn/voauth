"use client";
import Image from "next/image";
import FooterLogo from "@/public/assets/icons/logo-white.png";

export function Footer() {
  return (
    <footer className="bg-primary-01">
      <div className="max-w-[1440px] h-[86px] flex items-center px-[90px] mx-auto">
        <Image src={FooterLogo} width={97} height={22} alt="Voauth logo" />
        <span className="text-white ml-[100px] opacity-50">
          {`© ${new Date().getFullYear()} Volomn - All rights reserved`}
        </span>
      </div>
    </footer>
  );
}
