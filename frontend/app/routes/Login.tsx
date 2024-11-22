import { redirect } from "@remix-run/node";
import { Form } from "@remix-run/react";
import { Button } from "~/components/ui/button";
import { AuroraBackground } from "~/components/ui/aurora-background";
import { motion } from "framer-motion";

export async function action() {
  const CLIENT_ID = process.env.SPOTIFY_CLIENT_ID as string;
  const REDIRECT_URI = process.env.SPOTIFY_REDIRECT_URI as string;
  const SCOPE = "user-read-private user-read-email user-top-read";

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
    <AuroraBackground>
      <motion.div
        initial={{ opacity: 0.0, y: 40 }}
        whileInView={{ opacity: 1, y: 0 }}
        transition={{
          delay: 0.3,
          duration: 0.8,
          ease: "easeInOut",
        }}
        className="relative flex flex-col gap-4 items-center justify-center px-4"
      >
        <div className="text-3xl md:text-7xl font-bold dark:text-white text-center">
          Welcome to Onpu♪
        </div>
        <div className="font-extralight text-base md:text-4xl dark:text-neutral-200 py-4">
          Push the button!
        </div>
        <Form method="post" action="/login">
          <Button type="submit">Login with Spotify</Button>
        </Form>
      </motion.div>
    </AuroraBackground>
  );
}
