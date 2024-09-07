import React, { useState } from 'react'
import { Button } from '~/components/ui/button'
import X from '/x_logo.png'
import Instagram from '/instagram_logo.png'
import { ProfileProps } from '~/types/types'
import { useParams } from "@remix-run/react";

const Profile: React.FC<ProfileProps> = ({displayName, iconImage, xLink, instagramLink, myUserId}) => {
  const params = useParams()
  const pageUserId = params.user_id ? Number(params.user_id) : undefined;
  const [isFollowing, setIsFollowing] = useState(false);

  const handleFollow = async (): Promise<void> => {
    try {
      const response = await fetch(`http://localhost:8080/api/v1/follow/${myUserId}/${pageUserId}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
      });
  
      if (!response.ok) {
        const errorData: { message: string } = await response.json();
        throw new Error(errorData.message || `HTTPエラー! ステータス: ${response.status}`);
      }
  
      const data: { success: boolean } = await response.json();
      setIsFollowing(true);
      console.log("ユーザーのフォローに成功しました:", data);
    } catch (error) {
      console.error("ユーザーのフォローに失敗しました:", error instanceof Error ? error.message : String(error));
      // ここでエラー状態を設定し、ユーザーに表示することも検討してください
    }
  };
  

  return (
    <div className='flex flex-col gap-4'>
      <div className='flex justify-between'>
        <p className='flex justify-center items-center text-2xl'>Profile</p>
        <div className='pt-3'> 
          { 
            myUserId == pageUserId ? 
            <Button size="sm">Edit</Button> :
            <Button size="sm" onClick={handleFollow} disabled={isFollowing}>
              {isFollowing ? "Following" : "Follow"}
            </Button>
          }
        </div>
      </div>
      <div className='flex justify-between items-center px-2'>
        <div className='flex gap-4 justify-center items-center'>
          <img src={iconImage} className='w-20 h-20 rounded-full' alt={`${displayName}'s profile`} />
          <p className='flex justify-center items-center text-xl'>{displayName}</p>
        </div>
        <div className='flex gap-4'>
          <div className='flex gap-2 items-center'>
            <a href={xLink} target="_blank" rel="noopener noreferrer">
              <img src={X} className='w-6 h-fit' alt="X (Twitter) logo" />
            </a>
          </div>
          <div className='flex gap-2 items-center'>
            <a href={instagramLink} target="_blank" rel="noopener noreferrer">
              <img src={Instagram} className='w-6 h-fit' alt="Instagram logo" />
            </a>
          </div>
        </div>
      </div>
    </div>
  )
}

export { Profile }