import type { Route } from "./+types/home";

export function meta({ }: Route.MetaArgs) {
  return [
    { title: "About" },
  ];
}

export default function Home({ loaderData }: Route.ComponentProps) {
  return <div>about page</div>;
}
