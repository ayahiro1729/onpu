import { ActionFunctionArgs, redirect } from '@remix-run/node'

export const action = async ({ request }: ActionFunctionArgs) => {
  const formData = await request.formData()
  const followings_id = formData.get('followings_id')
  const follower_id = formData.get('follower_id')

  const response = await fetch(`http://backend:8080/api/v1/follow/${follower_id}/${followings_id}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  })

  if (!response.ok) {
    throw new Error('Failed to follow user')
  }

  return redirect(`/user/${followings_id}`)
}

export default function Follow() {

  return (
    <div>
      <p>Follow</p>
    </div>
  )
}