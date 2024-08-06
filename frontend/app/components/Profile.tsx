import React from 'react'
import { Button } from '~/components/ui/button'
import { Avatar, AvatarFallback, AvatarImage } from "~/components/ui/avatar"

const Profile = () => {
  return (
    <div className='flex flex-col gap-2'>
      <div className='flex justify-between'>
        <p className='flex justify-center items-center text-2xl'>Profile</p>
        <div className='pt-3'>
          <Button size="sm">Add a friend</Button>
        </div>
      </div>
      <div className='flex justify-between px-2'>
        <div className='flex gap-4'>
          <Avatar>
            <AvatarImage src="https://github.com/shadcn.png" />
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>
          <p className='flex justify-center items-center text-xl'>Name</p>
        </div>
        <div className='flex flex-col'>
          <a href='https://x.com'>X</a>
          <a href='https://www.instagram.com/'>Instagram</a>
        </div>
      </div>
    </div>
  )
}

export { Profile }