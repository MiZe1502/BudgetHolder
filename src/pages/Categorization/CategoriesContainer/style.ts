import {css} from 'emotion'

import colors from "../../../common/styles/colors";

export const maxHeight = 500;

export const Container = css`
  border: 1px solid ${colors.darkGray};
  border-radius: 4px;
  width: 100%;
  max-height: ${maxHeight}px;
  overflow: scroll;
`;
