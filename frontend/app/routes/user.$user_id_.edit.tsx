import React, { useEffect, useState } from "react";
import { Label } from "~/components/ui/label";
import { Input } from "~/components/ui/input";
import { cn } from "~/lib/utils";
import { Form, json, redirect, useLoaderData } from "@remix-run/react";
import { ActionFunctionArgs, LoaderFunctionArgs } from "@remix-run/node";
import { Header } from "~/components/Header";
import { UserInfo } from "~/types/types";

export const action = async ({
  params,
  request,
}: ActionFunctionArgs) => {
  const userId = params.user_id;
  const formData = await request.formData();
  const updates = Object.fromEntries(formData);

  const response = await fetch(`http://backend:8080/api/v1/user/${userId}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(updates),
  });

  if (!response.ok) {
    throw new Error('Failed to update music list');
  }

  return redirect(`/user/${userId}`);
};

export const loader = async ({ params }: LoaderFunctionArgs) => {
  const userId = params.user_id;
  const userResponse = await fetch(`http://backend:8080/api/v1/user/${userId}`);
  if (!userResponse.ok) {
    throw new Error (`Failed to fetch user data: ${userResponse.statusText}`)
  }
  const userData = await userResponse.json();

  const userInfo: UserInfo = {
    displayName: userData.user.DisplayName,
    iconImage: userData.user.IconImage,
    xLink: userData.user.XLink,
    instagramLink: userData.user.InstagramLink,
  };

  return json({ userInfo, userId });
};

export default function EditContact() {
  const { userInfo } = useLoaderData<typeof loader>();
  const [myUserId, setMyUserId] = useState<number | null>(null);

  useEffect(() => {
    const getMyUserId = async () => {
      try {
        const response = await fetch(`http://localhost:8080/api/v1/myuserid`, { credentials: "include" });
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const result = await response.json();
        setMyUserId(result.user_id);
      } catch (error) {
        console.error("There was a problem with the fetch operation:", error);
      }
    }
    getMyUserId();
  }, []);

  useEffect(() => {
    if (myUserId !== null) {
      console.log('myUserId:', myUserId);
    }
  }, [myUserId]);

  if (myUserId === null) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <Header
        myUserId={myUserId}
      />
      <div className="p-4 pt-20">
        <div className="max-w-md w-full mx-auto rounded-none md:rounded-2xl shadow-input bg-white dark:bg-black">
          <Form id="contact-form" method="post">
            <LabelInputContainer className="mb-4">
              <Label htmlFor="x_link">X URL</Label>
              <Input
                name="x_link"
                placeholder="https://x.com/example"
                type="text"
                id="x_link"
              />
            </LabelInputContainer>
            <LabelInputContainer className="mb-4">
              <Label htmlFor="instagram_link">Instagram URL</Label>
              <Input
                name="instagram_link"
                placeholder="https://instagram.com/example"
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