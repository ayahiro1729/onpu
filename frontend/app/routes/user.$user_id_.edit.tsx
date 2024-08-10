import React from "react";
import { Label } from "~/components/ui/label";
import { Input } from "~/components/ui/input";
import { cn } from "~/lib/utils";
import { Form, json, redirect, useLoaderData, useParams } from "@remix-run/react";
import { ActionFunctionArgs, LoaderFunctionArgs } from "@remix-run/node";

export const action = async ({
  params,
  request,
}: ActionFunctionArgs) => {
  const formData = await request.formData();
  const updates = Object.fromEntries(formData);
  await updateContact(params.userId, updates);
  return redirect(`https://localhost:8080/api/v1/user/${params.userId}`);
};

export const loader = async ({ params }: LoaderFunctionArgs) => {
  const response = await fetch(`https://localhost:8080/api/v1/user/${params.userId}`);
  const data = await response.json();

  return json({
    user_id: params.userId,
    xLink: data.x_link,
    instagramLink: data.instagram_link
  });
};

export default function EditContact() {
  const { user_id, xLink, instagramLink } = useLoaderData<typeof loader>();

  return (
    <div className="p-4 pt-20">
      <div className="max-w-md w-full mx-auto rounded-none md:rounded-2xl shadow-input bg-white dark:bg-black">
        <Form key={user_id} id="contact-form" method="post">
          <LabelInputContainer className="mb-4">
            <Label htmlFor="x_link">X URL</Label>
            <Input
              defaultValue={xLink}
              name="x_link"
              placeholder="@hoge"
              type="text"
              id="x_link"
            />
          </LabelInputContainer>
          <LabelInputContainer className="mb-4">
            <Label htmlFor="instagram_link">X URL</Label>
            <Input
              defaultValue={instagramLink}
              name="instagram_link"
              placeholder="@hoge"
              type="text"
              id="instagram_link"
            />
          </LabelInputContainer>
          <button
            className="bg-gradient-to-br relative group/btn from-black dark:from-zinc-900 dark:to-zinc-900 to-neutral-600 block dark:bg-zinc-800 w-full text-white rounded-md h-10 font-medium shadow-[0px_1px_0px_0px_#ffffff40_inset,0px_-1px_0px_0px_#ffffff40_inset] dark:shadow-[0px_1px_0px_0px_var(--zinc-800)_inset,0px_-1px_0px_0px_var(--zinc-800)_inset]"
            type="submit"
          >
            Update
            <BottomGradient />
          </button>
        </Form>
      </div>
    </div>
  );
}

const BottomGradient = () => {
  return (
    <>
      <span className="group-hover/btn:opacity-100 block transition duration-500 opacity-0 absolute h-px w-full -bottom-px inset-x-0 bg-gradient-to-r from-transparent via-cyan-500 to-transparent" />
      <span className="group-hover/btn:opacity-100 blur-sm block transition duration-500 opacity-0 absolute h-px w-1/2 mx-auto -bottom-px inset-x-10 bg-gradient-to-r from-transparent via-indigo-500 to-transparent" />
    </>
  );
};

const LabelInputContainer = ({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}) => {
  return (
    <div className={cn("flex flex-col space-y-2 w-full", className)}>
      {children}
    </div>
  );
};
function updateContact(userId: string | undefined, updates: { [k: string]: FormDataEntryValue; }) {
  throw new Error("Function not implemented.");
}

