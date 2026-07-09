import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function extractErrorMsg(error: unknown, defaultMsg: string = "Unknown error occurred"): string {
  if (!error) return defaultMsg;
  const errObj = error as Record<string, unknown>;
  
  if (errObj.response && typeof errObj.response === 'object') {
    const res = errObj.response as Record<string, unknown>;
    if (res.data && typeof res.data === 'object') {
      const data = res.data as Record<string, unknown>;
      if (typeof data.error === 'string') return data.error;
      if (typeof data.message === 'string') return data.message;
    }
  }
  
  if (typeof errObj.message === 'string') return errObj.message;
  return defaultMsg;
}
