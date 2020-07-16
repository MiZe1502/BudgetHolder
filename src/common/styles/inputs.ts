import { css } from 'emotion'

import colors from "./colors";

export const InvalidInput = css`
    border: 1px solid ${colors.redAttention} !important;
      
    &:focus {
        border: 1px solid ${colors.redAttention} !important;
        box-shadow: 0px 2px 1px -1px rgba(181,99,99,1) !important;
    }
`;