import Image from "next/image";
import Fingerprint from "@/public/assets/icons/fingerprint.png";
import { Capabilities, Elevate } from "@/components/home-images";
import ElevateImage from "@/public/assets/icons/elevate.svg";
export default function Home() {
  return (
    <>
      <header>
        <nav className="h-[90px] px-4 max-w-[1440px] mx-auto flex items-center">
          {/* <Logo /> */}
          <Image
            src="/assets/icons/logo.svg"
            width={100}
            height={18}
            alt="logo"
          />
        </nav>
      </header>
      <main>
        <section id="banner" className="h-[60vh] flex items-center relative">
          <div className="absolute top-0 w-full flex justify-center">
            <Image src={Fingerprint} alt="" />
          </div>
          <div className="flex flex-col gap-4 max-w-2xl mx-auto text-center relative">
            <h1 className="text-primary-01 text-5xl font-semibold font-secondary">
              Unlock the power of OAuth2 with Voauth
            </h1>
            <article className="text-shade-01 text-lg">
              The Voauth application is a web-based tool designed to show
              developers how to create their own OAuth provider platforms,
              similar to those of Google, Twitter, GitHub, and more.
            </article>
            <button className="bg-primary-01 px-8 py-5 rounded-lg text-white mt-[26px] w-fit mx-auto">
              Get Started
            </button>
          </div>
        </section>

        <section className="flex justify-between items-center px-4 max-w-[1440px] mx-auto py-[150px]">
          <div className="max-w-[562px]">
            <h2 className="text-primary-01 leading-[48px] text-[40px] font-semibold font-secondary">
              Elevate your note taking experience with Voauth
            </h2>
            <p className="text-shade-01 text-lg mt-4">
              Voauth is a notes application that allows users to register, make
              notes, and then easily share those notes with third-party apps
              using OAuth2.
            </p>
          </div>
          <Elevate />
        </section>

        <section className="px-4 max-w-[1440px] mx-auto">
          <div className="max-w-[562px]">
            <h2 className="text-primary-01 leading-[48px] text-[40px] font-semibold font-secondary">
              The Voauth application has its capabilities
            </h2>
            <p className="text-shade-01 text-lg mt-4">
              Voauth comes with various features that make it a flexible and
              powerful tool for users.
            </p>
          </div>
          <section className="flex justify-between px-4 max-w-[1440px] mx-auto py-[50px]">
            <Capabilities />

            <div className="flex flex-col gap-5 max-w-[520px]">
              <div className="p-4 rounded-lg bg-[#E7F4FF]">
                <h3 className="text-primary-01 font-semibold text-lg">
                  Effortless Registration and Easy Access
                </h3>
                <p className="mt-[10px] text-shade-01">
                  Creating an account on Voath is a straightforward and
                  user-friendly process and once registered, users can easily
                  and securely log in to their accounts using their email and
                  password.
                </p>
              </div>

              <div className="p-4 rounded-lg bg-accent-10">
                <h3 className="text-primary-01 font-semibold text-lg">
                  Easy Note Handling and Safe Sharing
                </h3>
                <p className="mt-[10px] text-shade-01">
                  Users can create, edit and manage notes using Voauth and can
                  confidently share their notes with others, knowing that their
                  content will remain secure and protected.
                </p>
              </div>

              <div className="p-4 rounded-lg bg-[#E7F4FF]">
                <h3 className="text-primary-01 font-semibold text-lg">
                  Craft Your Own OAuth App with Voauth
                </h3>
                <p className="mt-[10px] text-shade-01">
                  {`With Voauth, you can take control and build your very own OAuth
                application. Whether it's for enhancing your project's security,
                expanding your app's functionality, or exploring OAuth2
                technology`}
                </p>
              </div>

              <div className="p-4 rounded-lg bg-accent-10">
                <h3 className="text-primary-01 font-semibold text-lg">
                  {`Add 'Login with Voauth' to Your Apps`}
                </h3>
                <p className="mt-[10px] text-shade-01">
                  With Voauth, you can effortlessly integrate its authentication
                  system into your apps, making it simpler for users to sign in
                  and access their accounts.
                </p>
              </div>
            </div>
          </section>
        </section>
        <section className="flex justify-between max-w-[1440px] mx-auto py-[100px] px-[90px] bg-[#FBFBFB] border rounded-lg">
          <div className="max-w-[600px] flex flex-col gap-4 justify-start">
            <h2 className="font-semibold text-primary-01 text-[40px] font-secondary tracking-[-1.74px]">
              Tech Stack Used
            </h2>
            <p className="text-shade-01">
              {`At Voauth, we take pride in the advanced technology stack we've
            chosen to create a powerful and reliable platform.`}
            </p>
            <button className="bg-primary-01 px-8 py-5 rounded-lg text-white mt-2 w-fit">
              Get Started
            </button>
          </div>
        </section>
        <section className="flex justify-center max-w-[1440px] mx-auto py-[100px] px-4">
          <div className="flex flex-col items-center text-center gap-4">
            <h2 className="font-semibold text-primary-01 text-[40px] font-secondary tracking-[-1.74px]">
              Github Repository Link
            </h2>
            <p className="text-shade-01">
              Access and explore the source code, documentation, and project
              files of Voauth.
            </p>
            <button className="border-2 border-primary-01 px-8 py-5 rounded-lg text-primary-01 w-fit mt-2 font-medium">
              View project on Github
            </button>
          </div>
        </section>
      </main>

      {/* <section className="py-[100px] px-[90px] bg-[#E7F4FF4D] rounded-lg">
        <div className="max-w-[1440px] mx-auto">
          <h2 className="font-semibold text-primary-01 text-[40px] font-secondary tracking-[-1.74px]">
            Insights and Updates from Voauth
          </h2>

          <div className="grid grid-cols-3 gap-[54px] h-[450px] mt-8">
            <div className="border rounded-lg bg-slate-100 "></div>
            <div className="border rounded-lg bg-slate-100 "></div>
            <div className="border rounded-lg bg-slate-100 "></div>
          </div>
        </div>
      </section> */}

      <footer className="bg-primary-01">
        <div className="max-w-[1440px] h-[86px] flex items-center px-4 mx-auto">
          <Image
            src="/assets/icons/logo-white.svg"
            width={100}
            height={18}
            alt="logo"
          />

          <span className="text-white ml-[100px] opacity-50">
            {` © ${new Date().getFullYear()} Volomn - All rights reserved`}
          </span>
        </div>
      </footer>
    </>
  );
}
