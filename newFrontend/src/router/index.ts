import { createRouter, createWebHistory } from "vue-router";
const IndexView = () => import("../views/index.vue");
const AuthorView = () => import("../views/author.vue");
const GalleryView = () => import("../views/gallery.vue");
const NotFoundView = () => import("../views/404.vue");
const BooksView = () => import("../views/books.vue");
const BookView = () => import("../views/book.vue");
const BookManagerView = () => import("../views/admin/bookManager.vue");
const AdminView = () => import("../views/admin/admin.vue");
const ContactView = () => import("../views/contact.vue");
const SitemapView = () => import("../views/sitemap.vue");
const UpdateView = () => import("../views/update.vue");


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
       {
      path: "/:pathMatch(.*)*",
      name: "404",
      component: NotFoundView,
    },

  ],
});

export default router;
