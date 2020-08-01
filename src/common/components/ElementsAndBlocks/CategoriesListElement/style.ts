import {css} from "emotion";

export { Font312Black } from "../../../styles/fonts";
export { FlexVert } from "../../../styles/positioning";

import colors from "../../../styles/colors";

export const Delimeter = css`
  width: 20px;
  border-bottom: 2px solid ${colors.greenPrimary};
`

export const SingleLine = css`
    border-left: 2px solid ${colors.greenPrimary};
    padding-left: 8px;
    padding-top: 2px;
`;