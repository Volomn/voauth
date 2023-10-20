import localFont from "next/font/local";

const sora = localFont({
  src: "../fonts/Sora/Sora-VariableFont_wght.ttf",
  variable: "--font-secondary",
});

const montserrat = localFont({
  src: "../fonts/Montserrat/Montserrat-VariableFont_wght.ttf",
  variable: "--font-primary",
});

export { sora, montserrat };
