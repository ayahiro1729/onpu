import React, { useState, useEffect } from 'react'
import { Button } from '~/components/ui/button'
import X from '/x_logo.png'
import Instagram from '/instagram_logo.png'
import { ProfileProps } from '~/types/types'
import { Form, useParams } from "@remix-run/react";

const Profile: React.FC<ProfileProps> = ({displayName, iconImage, xLink, instagramLink, myUserId}) => {
  const params = useParams()
  const pageUserId = params.user_id
  const [isFollowed, setIsFollowed] = useState<boolean>(false)

  useEffect(() => {
    const followJudge = async () => {
      const response = await fetch(`http://localhost:8080/api/v1/is_following/${myUserId}/${pageUserId}`)
      if (!response.ok) {
        throw new Error (`Failed to fetch follower data: ${response.statusText}`)
      }
      const data = await response.json()
      setIsFollowed(data.is_following)
    }
    followJudge()
  }, [isFollowed])

  return (
    <div className='flex flex-col gap-4'>
      <div className='flex justify-between'>
        <p className='flex justify-center items-center text-2xl'>Profile</p>
        <div className='pt-3'>
          {
            myUserId == pageUserId ?
            <Form action="edit">
              <Button size="sm">Edit</Button>
            </Form> :
            isFollowed ?
            <Form method='post' action='/unfollow'>
              <input type='hidden' name='followings_id' value={pageUserId}/>
              { myUserId && <input type='hidden' name='follower_id' value={myUserId}/> }
              <Button type='submit' name='_action' aria-label='delete' value='delete' onClick={() => location.reload()}>Unfollow</Button>
            </Form> :
            <Form method='post' action='/follow'>
              <input type='hidden' name='followings_id' value={pageUserId}/>
              { myUserId && <input type='hidden' name='follower_id' value={myUserId}/> }
              <Button type='submit' name='_action' aria-label='post' value='post' variant="outline" onClick={() => location.reload()}>Follow</Button>
            </Form>
          }
        </div>
      </div>
      <div className='flex justify-between items-center px-2'>
        <div className='flex gap-4 justify-center items-center'>
          <img src={iconImage} className='w-20 h-20 rounded-full'/>
          <p className='flex justify-center items-center text-xl'>{displayName}</p>
        </div>
        <div className='flex gap-4'>
          <div className='flex gap-2 items-center'>
            <a href={xLink}>
              <img src={X} className='w-6 h-fit'/>
            </a>
          </div>
          <div className='flex gap-2 items-center'>
            <a href={instagramLink}>
              <img src={Instagram} className='w-6 h-fit'/>
            </a>
          </div>
        </div>
      </div>
    </div>
  )
}

export { Profile }