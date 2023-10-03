import Link from 'next/link';
import { useRouter } from 'next/router';
import { Button } from 'nextra/components';

function Page404() {
  const { asPath } = useRouter();
  let message = 'This page could not be found.';

  if (asPath.includes('/u/')) {
    message = 'Not Found';
  }

  return (
    <div className="h-screen w-screen bg-white grid place-items-center">
      <div className="flex flex-col items-center space-y-10">
        <div className="text-black/75 flex justify-center items-center">
          <h2 className="text-2xl font-medium">404</h2>
          <span className="w-[1px] bg-black/30 h-12 mx-8"></span>
          <p>{message}</p>
        </div>
        <Link href="/">
          <Button>Back To Home</Button>
        </Link>
      </div>
    </div>
  );
}

export default Page404;
