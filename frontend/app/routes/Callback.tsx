import { json, LoaderFunction } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";

export const loader: LoaderFunction = async ({ request }): Promise<Response> => {
  const url = new URL(request.url);
  const code = url.searchParams.get("code");
  const error = url.searchParams.get("error");

  console.log("url:", url);
  console.log("code:", code);

  if (error) {
    return json({ error });
  }

  if (!code) {
    return json({ error: "No code provided" });
  }

  try {
    // GoのAPIにcodeを送信
    const response = await fetch("http://localhost:8080/api/v1/user", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ code }),
    });

    if (!response.ok) {
      throw new Error("Failed to exchange code for token");
    }

    const data = await response.json();

    // ここでアクセストークンを安全に保存する処理を行う
    // 例: セッションに保存したり、HTTPSのみのクッキーに保存したりします

    return json({ success: true });
  } catch (error) {
    console.error("Error exchanging code for token:", error);
    return json({ error: "Failed to exchange code for token" });
  }
};

export default function Callback() {
  const data = useLoaderData<typeof loader>();

  if (data.error) {
    return <div className="p-4 pt-20">Error: {data.error}</div>;
  }

  if (data.success) {
    return <div className="p-4 pt-20">Successfully authenticated with Spotify!</div>;
  }

  return <div className="p-4 pt-20">Processing...</div>;
}