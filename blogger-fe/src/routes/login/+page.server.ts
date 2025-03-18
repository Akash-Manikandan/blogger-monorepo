import type { PageServerLoad, Actions } from "./$types";
import { fail } from "@sveltejs/kit";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { formSchema } from "./schema";
import { userClient } from "$lib/grpc/users";

export const load: PageServerLoad = async () => {
 return {
  form: await superValidate(zod(formSchema)),
 };
};
 
export const actions: Actions = {
 default: async (event) => {
  const form = await superValidate(event, zod(formSchema));
  const data = await userClient.login(form.data);
  console.log(data);
  if (!form.valid) {
   return fail(400, {
    form,
   });
  }
  return {
   form,
  };
 },
};