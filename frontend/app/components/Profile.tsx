import React, { useState } from 'react'
import { Button } from '~/components/ui/button'
import X from '/x_logo.png'
import Instagram from '/instagram_logo.png'
import { Link } from '@remix-run/react'
import { ProfileProps } from '~/types/types'

const Profile: React.FC<ProfileProps> = ({displayName, iconImage, xLink, instagramLink}) => {
  const [isFollowed, setIsFollowed] = useState<boolean>(false)

  const handleFollow = () => {
    setIsFollowed(!isFollowed)
  }

  return (
    <div className='flex flex-col gap-4'>
      <div className='flex justify-between'>
        <p className='flex justify-center items-center text-2xl'>Profile</p>
        <div className='pt-3'> 
          { isFollowed ? <Button size="sm" onClick={handleFollow}>Unfollow</Button> : <Button size="sm" onClick={handleFollow} variant="outline">Follow</Button> }
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