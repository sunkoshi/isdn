import { createTRPCReact } from "@trpc/react";
import type { AppRouter } from "@todo/app/src/routes";

export const trpc = createTRPCReact<AppRouter>();
