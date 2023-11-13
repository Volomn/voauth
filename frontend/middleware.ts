export { default } from "next-auth/middleware";

export const config = {
  matcher: ["/dashboard", "/notes", "/favourites", "/bin", "/archive"],
};
