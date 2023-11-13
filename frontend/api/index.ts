import axios from "axios";

export const axiosInstance = axios.create({
  baseURL: process.env.APP_BASE_URL,
});

export const _axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_APP_BASE_URL,
});
