import React from 'react'
import { Button } from '~/components/ui/button'
import { Avatar, AvatarFallback, AvatarImage } from "~/components/ui/avatar"
import X from '/x_logo.png'
import Instagram from '/instagram_logo.png'

const Profile = () => {
  return (
    <div className='flex flex-col gap-4'>
      <div className='flex justify-between'>
        <p className='flex justify-center items-center text-2xl'>Profile</p>
        <div className='pt-3'>
          <Button size="sm">Add a friend</Button>
        </div>
      </div>
      <div className='flex justify-between items-center px-2'>
        <div className='flex gap-4 justify-center items-center'>
          <Avatar>
            <AvatarImage src="https://github.com/shadcn.png" />
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>
          <p className='flex justify-center items-center text-xl'>Name</p>
        </div>
        <div className='flex flex-col gap-2'>
          <div className='flex gap-2'>
            <img src={X} className='w-5 h-auto'/>
            <a href='https://x.com'>@hoge</a>
          </div>
          <div className='flex gap-2'>
            <img src={Instagram} className='w-5 h-auto'/>
            <a href='https://www.instagram.com/'>@hoge</a>
          </div>
        </div>
      </div>
    </div>
  )
}

export { Profile }