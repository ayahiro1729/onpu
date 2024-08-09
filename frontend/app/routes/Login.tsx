import { redirect } from "@remix-run/node";
import { Form } from "@remix-run/react";
import { Button } from "~/components/ui/button";

export async function action() {
  const CLIENT_ID = process.env.SPOTIFY_CLIENT_ID as string;
  const REDIRECT_URI = process.env.SPOTIFY_REDIRECT_URI as string;
  const SCOPE = "user-read-private user-read-email";

  const params = new URLSearchParams({
    response_type: "code",
    client_id: CLIENT_ID,
    scope: SCOPE,
    redirect_uri: REDIRECT_URI,
  });

  return redirect(`https://accounts.spotify.com/authorize?${params}`);
}

export default function Login() {
  return (
    <div className="p-4 pt-20">
      <Form method="post" action="/login">
        <Button type="submit">Login with Spotify</Button>
      </Form>
    </div>
  );
}