import { PrismaClient } from "@prisma/client";
import { initTRPC } from "@trpc/server";
import { number, z } from "zod";
const prisma = new PrismaClient();
export const t = initTRPC.create();
export const appRouter = t.router({
  getTodos: t.procedure
    .input(
      z.object({
        limit: z.number().default(10),
        offset: z.number().default(0),
      })
    )
    .query(async (req) => {
      return prisma.todos.findMany({
        take: req.input.limit,
        skip: req.input.offset,
      });
    }),
  getTodoById: t.procedure
    .input(z.object({ id: z.number() }))
    .query(async (req) => {
      return prisma.todos.findUnique({ where: { id: req.input.id } });
    }),
  createTodo: t.procedure
    .input(
      z.object({
        title: z.string().min(5),
        description: z.string().optional(),
      })
    )
    .mutation(async (req) => {
      const inputs = req.input;
      try {
        const todo = await prisma.todos.create({
          data: { title: req.input.title, description: req.input.description },
        });
        return { data: todo, message: "todo created" };
      } catch (err) {
        throw err;
      }
    }),
});

export type AppRouter = typeof appRouter;
