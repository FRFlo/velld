export interface Pagination {
  total: number;
  page: number;
  limit: number;
  total_pages: number;
}

export interface Base<T> {
  data: T;
  message: string;
  pagination?: Pagination;
}