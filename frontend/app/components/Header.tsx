import React, { useState, useEffect } from 'react'
import logo from '/OnpuLogo.jpg'
import { UserSearch } from './UserSearch'
import { Link } from '@remix-run/react'
import { HeaderProps } from '~/types/types'

export const Header :React.FC<HeaderProps> = ({ myUserId }) => {
  const [icon, setIcon] = useState<string>('')

  useEffect(() => {
    const getIcon = async () => {
      try {
        const response = await fetch(`http://localhost:8080/api/v1/user/${myUserId}`)
        if (!response.ok) {
          throw new Error(`Failed to fetch user data: ${response.statusText}`)
        }
        const data = await response.json()
        setIcon(data.user.icon_image)
      }
      catch (error) {
        console.error("Cannot get icon image", error)
      }
    }
    getIcon()
  }, [])

  return (
    <div className="p-2 flex justify-between items-center fixed top-0 left-0 right-0 bg-white border-b-2 border-current z-[9999]">
      <Link to={`/user/${myUserId}`}>
        <div className='flex flex-column gap-1'>
          <img src={logo} className="w-12 shadow rounded-full max-w-full h-auto align-middle border-none" />
          <p className='flex justify-center items-center'>Onpu</p>
        </div>
      </Link>
      <div className='flex flex-column gap-3'>
        <UserSearch />
        <Link to={`/user/${myUserId}/edit`}>
          <img src={icon} className='w-8 h-8 rounded-full'/>
        </Link>
      </div>
    </div>
  )
}