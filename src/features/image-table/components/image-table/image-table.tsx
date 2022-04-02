import ImageTableRow from "@features/image-table/components/image-table/image-table.row";
import { IMAGE_TABLE_HEADERS } from "@features/image-table/image-table.constants";

import { Table } from "@mantine/core";

import type { Image } from "@graphql/generated/codegen.generated";

export type ImageTableProps = {
  data: Image[];
};

const ImageTable = ({ data }: ImageTableProps): JSX.Element => {
  const rows = data.map((image, index: number) => (
    <ImageTableRow
      image={image}
      key={index}
    />
  ));

  return (
    <Table
      striped
      highlightOnHover
    >
      <thead>
        <tr>
          {IMAGE_TABLE_HEADERS.map((header, index) => (
            <th key={index}>{header}</th>
          ))}
        </tr>
      </thead>
      <tbody>{rows}</tbody>
    </Table>
  );
};

export default ImageTable;
