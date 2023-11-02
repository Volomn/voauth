"use client";
import Image from "next/image";
import NavLogo from "@/public/assets/icons/logo.png";

export function Navbar() {
  return (
    <header>
      <nav className="h-[90px] px-[90px] max-w-[1440px] mx-auto flex items-center">
        <Image src={NavLogo} width={97} height={22} alt="Voauth logo" />
      </nav>
    </header>
  );
}
