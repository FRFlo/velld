import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";

interface CustomPagination {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
}

export function CustomPagination({
  currentPage,
  totalPages,
  onPageChange,
}: CustomPagination) {
  const handlePageChange = (page: number) => {
    if (page >= 1 && page <= totalPages) {
      onPageChange(page);
    }
  };

  const getVisiblePages = () => {
    const visiblePages = [];
    const showEllipsisStart = currentPage > 3;
    const showEllipsisEnd = currentPage < totalPages - 2;

    if (totalPages <= 5) {
      for (let i = 1; i <= totalPages; i++) {
        visiblePages.push(i);
      }
    } else {
      visiblePages.push(1);

      if (showEllipsisStart) {
        visiblePages.push('ellipsis1');
      }

      for (let i = Math.max(2, currentPage - 1); 
           i <= Math.min(currentPage + 1, totalPages - 1); i++) {
        if (!visiblePages.includes(i)) {
          visiblePages.push(i);
        }
      }

      if (showEllipsisEnd) {
        visiblePages.push('ellipsis2');
      }

      if (!visiblePages.includes(totalPages)) {
        visiblePages.push(totalPages);
      }
    }

    return visiblePages;
  };

  return (
    <Pagination>
      <PaginationContent>
        <PaginationItem>
          <PaginationPrevious
            href="#"
            onClick={(e) => {
              e.preventDefault();
              handlePageChange(currentPage - 1);
            }}
            className={currentPage <= 1 ? 'pointer-events-none opacity-50' : ''}
          />
        </PaginationItem>

        {getVisiblePages().map((pageNum, index) => (
          <PaginationItem key={pageNum + '-' + index}>
            {typeof pageNum === 'string' ? (
              <span className="px-4">...</span>
            ) : (
              <PaginationLink
                href="#"
                isActive={pageNum === currentPage}
                onClick={(e) => {
                  e.preventDefault();
                  handlePageChange(Number(pageNum));
                }}
              >
                {pageNum}
              </PaginationLink>
            )}
          </PaginationItem>
        ))}

        <PaginationItem>
          <PaginationNext
            href="#"
            onClick={(e) => {
              e.preventDefault();
              handlePageChange(currentPage + 1);
            }}
            className={currentPage >= totalPages ? 'pointer-events-none opacity-50' : ''}
          />
        </PaginationItem>
      </PaginationContent>
    </Pagination>
  );
}
