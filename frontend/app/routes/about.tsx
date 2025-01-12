import type { Route } from "./+types/home";

export function meta({ }: Route.MetaArgs) {
  return [
    { title: "About" },
  ];
}


export async function clientLoader({
  params,
}: Route.ClientLoaderArgs) {
  // sleep for 2s
  await new Promise((resolve) => setTimeout(resolve, 2000));
  return {}
}


export default function Home({ loaderData }: Route.ComponentProps) {
  return <div>about page</div>;
}
