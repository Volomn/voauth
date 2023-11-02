import { axiosInstance } from "@/api";
import NextAuth from "next-auth/next";
import CredentialsProvider from "next-auth/providers/credentials";

const handler = NextAuth({
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        userName: { label: "Username", type: "text", placeholder: "jsmith" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials, req) {
        // Add logic here to look up the user from the credentials supplied
        // const user = { id: "1", name: "J Smith", email: "jsmith@example.com" };
        try {
          const user = await axiosInstance.post("/auth/", {
            userName: credentials?.userName,
            password: credentials?.password,
          });
          return user.data;
        } catch (error) {
          console.error(error);
          return null;
        }

        // console.log({ credentials, req });
        // if (user) {
        //   // Any object returned will be saved in `user` property of the JWT
        //   return user.data;
        // } else {
        //   // If you return null then an error will be displayed advising the user to check their details.
        //   return null;

        //   // You can also Reject this callback with an Error thus the user will be sent to the error page with the error message as a query parameter
        // }
      },
    }),
  ],
  pages: {
    signIn: "/login",
    signOut: "/register",
  },
});

export { handler as GET, handler as POST };
