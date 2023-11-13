import { ReactNode } from "react";
import { Box } from "@mantine/core";
import Image from "next/image";
import CloudImage from "@/public/assets/cloud.png";
import Link from "next/link";
import NavLogo from "@/public/assets/icons/logo.png";

export default function AuthLayout({ children }: { children: ReactNode }) {
  return (
    <main className="min-h-[100dvh] w-full flex items-center justify-center p-4">
      <section className="w-full max-w-[500px] border shadow-md rounded-md relative bg-white">
        <Box className="display-none sm:display-block absolute -left-[40%] -top-[25%]">
          <Image src={CloudImage} alt="" />
        </Box>
        <Box className="display-none sm:display-block absolute -right-[10%] -bottom-[25%]">
          <Image src={CloudImage} alt="" />
        </Box>
        <section className="p-6 flex flex-col gap-8 bg-white relative">
          <Link href="/">
            <Image
              src={NavLogo}
              width={97}
              height={22}
              alt="Voauth logo"
              className="cursor-pointer"
            />
          </Link>

          {children}
        </section>
      </section>
    </main>
  );
}
