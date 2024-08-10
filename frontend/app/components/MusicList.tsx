import React from "react";
import { Carousel, Card } from "~/components/ui/apple-cards-carousel";
import { Loader2 } from "lucide-react"
 
import { Button } from "~/components/ui/button"
import { json, LoaderFunction } from '@remix-run/node';
import { useLoaderData } from '@remix-run/react';
import { useParams } from "@remix-run/react";

type Music = {
  music_id: number,
  name: string,
  image: string,
  artist_name: string,
  spotify_link: string
};

type MusicCard = {
  src: string;
  title: string;
  category: string;
  content: string;
};

export const loader: LoaderFunction = async () => {
  const { user_id } = useParams();
  const response = await fetch(`https://localhost:8080/api/v1/music/${user_id}`);
  const data = await response.json();

  const musicList = data.music_list;
  const musicCardList = musicList.map((music: Music) => {
    return {
      src: music.image,
      title: music.name,
      category: music.artist_name,
      content: music.spotify_link
    };
  });

  return json({ musicCardList });
};

export function MusicList() {
  // const { musicCardList } = useLoaderData<typeof loader>();
  const musics = musicCardList.map((card: MusicCard, index: number) => (
    <a href={card.content}>
      <Card key={card.src} card={card} index={index} />
    </a>
  ));

  return (
    <div className="flex flex-col gap-4">
      <div className="flex justify-between items-center">
        <p className='flex items-center text-2xl'>Favorite Music</p>
        <Button className="px-2 py-1 bg-[#1ED760]">
          <Loader2 className="mr-2 h-4 w-4 animate-spin" />
            <p>Update</p>
        </Button>
      </div>
      <Carousel items={musics} />
    </div>
  );
}

const musicCardList: MusicCard[] = [
  {
    src: "/OnpuLogo.jpg",
    title: "Onpu",
    category: "Onpu",
    content: "https://open.spotify.com/track/4uLU6hMCjMI75M1A2tKUQC"
  },
  {
    src: "/OnpuLogo.jpg",
    title: "Onpu",
    category: "Onpu",
    content: "https://open.spotify.com/track/4uLU6hMCjMI75M1A2tKUQC"
  },
  {
    src: "/OnpuLogo.jpg",
    title: "Onpu",
    category: "Onpu",
    content: "https://open.spotify.com/track/4uLU6hMCjMI75M1A2tKUQC"
  },
  {
    src: "/OnpuLogo.jpg",
    title: "Onpu",
    category: "Onpu",
    content: "https://open.spotify.com/track/4uLU6hMCjMI75M1A2tKUQC"
  }
];