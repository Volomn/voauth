"use client";
import { signIn } from "next-auth/react";
export function GetStartedButton() {
  return (
    <button
      onClick={() => signIn()}
      className="bg-primary-01 px-8 py-5 rounded-lg text-white mt-[26px] w-fit mx-auto"
    >
      Get Started
    </button>
  );
}
